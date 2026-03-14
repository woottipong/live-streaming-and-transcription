package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iwoody/realtime-streaming/backend/internal/livekit"
	"github.com/iwoody/realtime-streaming/backend/internal/model"
	"github.com/iwoody/realtime-streaming/backend/internal/repository"
)

type SessionHandler struct {
	repo        repository.SessionRepository
	roomClient  livekit.RoomClient
	tokenClient livekit.TokenClient
}

const liveKitCleanupTimeout = 5 * time.Second

type createSessionRequest struct {
	Title           string  `json:"title"`
	Description     *string `json:"description"`
	SourceType      string  `json:"source_type"`
	Language        string  `json:"language"`
	ASRProvider     string  `json:"asr_provider"`
	SubtitleEnabled bool    `json:"subtitle_enabled"`
}

type updateSessionRequest struct {
	Title           patchStringField `json:"title"`
	Description     patchStringField `json:"description"`
	SourceType      patchStringField `json:"source_type"`
	Language        patchStringField `json:"language"`
	ASRProvider     patchStringField `json:"asr_provider"`
	SubtitleEnabled *bool            `json:"subtitle_enabled"`
}

type createSessionTokenRequest struct {
	Identity *string `json:"identity"`
	Name     *string `json:"name"`
}

type createSessionTokenResponse struct {
	Token    string `json:"token"`
	Identity string `json:"identity"`
	RoomName string `json:"room_name"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type patchStringField struct {
	Set   bool
	Value *string
}

func (f *patchStringField) UnmarshalJSON(data []byte) error {
	f.Set = true

	if string(data) == "null" {
		f.Value = nil
		return nil
	}

	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	f.Value = &value
	return nil
}

func NewSessionHandler(repo repository.SessionRepository, roomClient livekit.RoomClient, tokenClient livekit.TokenClient) *SessionHandler {
	return &SessionHandler{
		repo:        repo,
		roomClient:  roomClient,
		tokenClient: tokenClient,
	}
}

func (h *SessionHandler) Register(app *fiber.App) {
	api := app.Group("/api/sessions")
	api.Post("/", h.Create)
	api.Get("/:id", h.GetByID)
	api.Patch("/:id", h.Update)
	api.Post("/:id/token/publisher", h.CreatePublisherToken)
	api.Post("/:id/token/viewer", h.CreateViewerToken)
}

func (h *SessionHandler) createSessionWithLiveKit(ctx context.Context, params repository.CreateSessionParams) (model.Session, error) {
	if err := h.roomClient.CreateRoom(ctx, params.RoomName); err != nil {
		return model.Session{}, err
	}

	session, err := h.repo.Create(ctx, params)
	if err != nil {
		cleanupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), liveKitCleanupTimeout)
		defer cancel()

		if cleanupErr := h.roomClient.DeleteRoom(cleanupCtx, params.RoomName); cleanupErr != nil {
			return model.Session{}, fmt.Errorf("create session: %w (cleanup livekit room: %v)", err, cleanupErr)
		}

		return model.Session{}, err
	}

	return session, nil
}

func (h *SessionHandler) Create(c *fiber.Ctx) error {
	var req createSessionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: "invalid request body"})
	}

	params, err := req.toCreateParams()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: err.Error()})
	}

	session, err := h.createSessionWithLiveKit(c.UserContext(), params)
	if err != nil {
		if errors.Is(err, livekit.ErrCreateRoomFailed) {
			return c.Status(fiber.StatusServiceUnavailable).JSON(errorResponse{Error: "failed to create LiveKit room"})
		}
		if errors.Is(err, repository.ErrDuplicateRoomName) {
			return c.Status(fiber.StatusConflict).JSON(errorResponse{Error: "room name already exists, please try again"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: "failed to create session"})
	}

	return c.Status(fiber.StatusCreated).JSON(session)
}

func (h *SessionHandler) GetByID(c *fiber.Ctx) error {
	sessionID, err := parseSessionID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: err.Error()})
	}

	session, err := h.repo.GetByID(c.UserContext(), sessionID)
	if err != nil {
		if errors.Is(err, repository.ErrSessionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse{Error: "session not found"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: "failed to get session"})
	}

	return c.JSON(session)
}

func (h *SessionHandler) Update(c *fiber.Ctx) error {
	sessionID, err := parseSessionID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: err.Error()})
	}

	var req updateSessionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: "invalid request body"})
	}

	params, err := req.toUpdateParams()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: err.Error()})
	}

	session, err := h.repo.Update(c.UserContext(), sessionID, params)
	if err != nil {
		if errors.Is(err, repository.ErrSessionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse{Error: "session not found"})
		}
		if errors.Is(err, repository.ErrConstraintViolated) {
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: "invalid field value"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: "failed to update session"})
	}

	return c.JSON(session)
}

func (h *SessionHandler) CreatePublisherToken(c *fiber.Ctx) error {
	return h.createToken(c, "publisher")
}

func (h *SessionHandler) CreateViewerToken(c *fiber.Ctx) error {
	return h.createToken(c, "viewer")
}

func (h *SessionHandler) createToken(c *fiber.Ctx, role string) error {
	sessionID, err := parseSessionID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: err.Error()})
	}

	var req createSessionTokenRequest
	if len(c.Body()) > 0 {
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: "invalid request body"})
		}
	}

	session, err := h.repo.GetByID(c.UserContext(), sessionID)
	if err != nil {
		if errors.Is(err, repository.ErrSessionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse{Error: "session not found"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: "failed to get session"})
	}

	identity, err := req.identity(role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: err.Error()})
	}
	name, err := req.name(identity)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: err.Error()})
	}

	var token string
	switch role {
	case "publisher":
		token, err = h.tokenClient.CreatePublisherToken(session.RoomName, identity, name)
	case "viewer":
		token, err = h.tokenClient.CreateViewerToken(session.RoomName, identity, name)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: "invalid token role"})
	}
	if err != nil {
		if errors.Is(err, livekit.ErrCreateTokenFailed) {
			return c.Status(fiber.StatusServiceUnavailable).JSON(errorResponse{Error: "failed to create LiveKit token"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: "failed to create token"})
	}

	return c.JSON(createSessionTokenResponse{
		Token:    token,
		Identity: identity,
		RoomName: session.RoomName,
	})
}

func (r createSessionRequest) toCreateParams() (repository.CreateSessionParams, error) {
	title := strings.TrimSpace(r.Title)
	if title == "" {
		return repository.CreateSessionParams{}, errors.New("title is required")
	}

	sourceType, err := normalizeSourceType(r.SourceType)
	if err != nil {
		return repository.CreateSessionParams{}, err
	}

	language := strings.TrimSpace(r.Language)
	if language == "" {
		return repository.CreateSessionParams{}, errors.New("language is required")
	}

	asrProvider := strings.TrimSpace(r.ASRProvider)
	if asrProvider == "" {
		return repository.CreateSessionParams{}, errors.New("asr_provider is required")
	}

	return repository.CreateSessionParams{
		Title:           title,
		Description:     normalizeOptionalString(r.Description),
		RoomName:        generateRoomName(title),
		SourceType:      sourceType,
		Language:        language,
		ASRProvider:     asrProvider,
		SubtitleEnabled: r.SubtitleEnabled,
	}, nil
}

func (r updateSessionRequest) toUpdateParams() (repository.UpdateSessionParams, error) {
	params := repository.UpdateSessionParams{
		SubtitleEnabled: r.SubtitleEnabled,
	}

	if r.Title.Set {
		title, err := normalizeRequiredPatchString("title", r.Title.Value)
		if err != nil {
			return repository.UpdateSessionParams{}, err
		}
		params.Title = title
	}

	if r.Description.Set {
		if r.Description.Value == nil {
			params.ClearDescription = true
		} else {
			params.Description = normalizeOptionalString(r.Description.Value)
		}
	}

	if r.Language.Set {
		language, err := normalizeRequiredPatchString("language", r.Language.Value)
		if err != nil {
			return repository.UpdateSessionParams{}, err
		}
		params.Language = language
	}

	if r.ASRProvider.Set {
		asrProvider, err := normalizeRequiredPatchString("asr_provider", r.ASRProvider.Value)
		if err != nil {
			return repository.UpdateSessionParams{}, err
		}
		params.ASRProvider = asrProvider
	}

	if r.SourceType.Set {
		if r.SourceType.Value == nil {
			return repository.UpdateSessionParams{}, errors.New("source_type is required")
		}

		sourceType, err := normalizeSourceType(*r.SourceType.Value)
		if err != nil {
			return repository.UpdateSessionParams{}, err
		}
		params.SourceType = &sourceType
	}

	hasUpdate := params.Title != nil ||
		params.Description != nil ||
		params.ClearDescription ||
		params.Language != nil ||
		params.SourceType != nil ||
		params.ASRProvider != nil ||
		params.SubtitleEnabled != nil
	if !hasUpdate {
		return repository.UpdateSessionParams{}, errors.New("at least one field must be provided")
	}

	return params, nil
}

func normalizeSourceType(value string) (string, error) {
	sourceType := strings.TrimSpace(value)
	if sourceType == "" {
		return "", errors.New("source_type is required")
	}

	switch sourceType {
	case "browser", "rtmp":
		return sourceType, nil
	default:
		return "", errors.New("source_type must be one of: browser, rtmp")
	}
}

func normalizeOptionalString(value *string) *string {
	if value == nil {
		return nil
	}

	trimmed := strings.TrimSpace(*value)
	return &trimmed
}

func (r createSessionTokenRequest) identity(role string) (string, error) {
	if r.Identity == nil {
		return fmt.Sprintf("%s-%s", role, uuid.NewString()), nil
	}

	identity := strings.TrimSpace(*r.Identity)
	if identity == "" {
		return "", errors.New("identity is required")
	}

	return identity, nil
}

func (r createSessionTokenRequest) name(identity string) (string, error) {
	if r.Name == nil {
		return identity, nil
	}

	name := strings.TrimSpace(*r.Name)
	if name == "" {
		return "", errors.New("name is required")
	}

	return name, nil
}

func normalizeRequiredPatchString(name string, value *string) (*string, error) {
	trimmed := normalizeOptionalString(value)
	if trimmed == nil || *trimmed == "" {
		return nil, fmt.Errorf("%s is required", name)
	}

	return trimmed, nil
}

func parseSessionID(value string) (uuid.UUID, error) {
	sessionID, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, errors.New("invalid session id")
	}

	return sessionID, nil
}

func generateRoomName(title string) string {
	slugParts := strings.Fields(strings.ToLower(title))
	slug := strings.Join(slugParts, "-")
	if slug == "" {
		slug = "session"
	}

	return fmt.Sprintf("%s-%s", slug, uuid.NewString()[:8])
}
