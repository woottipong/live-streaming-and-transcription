package handler

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	livekitproto "github.com/livekit/protocol/livekit"

	"github.com/iwoody/realtime-streaming/backend/internal/livekit"
	"github.com/iwoody/realtime-streaming/backend/internal/repository"
)

type WebhookHandler struct {
	repo            repository.SessionRepository
	webhookVerifier livekit.WebhookVerifier
}

type webhookResponse struct {
	Status       string `json:"status"`
	StreamStatus string `json:"stream_status,omitempty"`
}

func NewWebhookHandler(repo repository.SessionRepository, webhookVerifier livekit.WebhookVerifier) *WebhookHandler {
	return &WebhookHandler{
		repo:            repo,
		webhookVerifier: webhookVerifier,
	}
}

func (h *WebhookHandler) Register(app *fiber.App) {
	app.Post("/internal/livekit/webhook", h.Receive)
}

func (h *WebhookHandler) Receive(c *fiber.Ctx) error {
	event, err := h.webhookVerifier.VerifyWebhookEvent(c.Get("Authorization"), c.Body())
	if err != nil {
		if errIsWebhookInvalid(err) {
			log.Printf("livekit webhook rejected: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse{Error: "invalid webhook signature"})
		}

		log.Printf("livekit webhook parse failed: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: "invalid webhook payload"})
	}

	roomName := webhookRoomName(event)
	streamStatus, shouldUpdate := streamStatusForEvent(event)
	participantIdentity := ""
	if participant := event.GetParticipant(); participant != nil {
		participantIdentity = participant.GetIdentity()
	}

	log.Printf("received livekit webhook event=%s room=%s participant=%s", event.GetEvent(), roomName, participantIdentity)

	if !shouldUpdate {
		return c.JSON(webhookResponse{Status: "ignored"})
	}

	if roomName == "" {
		log.Printf("livekit webhook ignored due to missing room name event=%s", event.GetEvent())
		return c.JSON(webhookResponse{Status: "ignored"})
	}

	session, err := h.repo.UpdateStreamStatusByRoomName(c.UserContext(), roomName, streamStatus)
	if err != nil {
		if errors.Is(err, repository.ErrSessionNotFound) {
			log.Printf("livekit webhook ignored because session was not found room=%s event=%s", roomName, event.GetEvent())
			return c.JSON(webhookResponse{Status: "ignored"})
		}

		log.Printf("livekit webhook failed to update session room=%s event=%s err=%v", roomName, event.GetEvent(), err)
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: "failed to update session"})
	}

	log.Printf("livekit webhook updated session room=%s stream_status=%s event=%s", roomName, session.StreamStatus, event.GetEvent())

	return c.JSON(webhookResponse{
		Status:       "ok",
		StreamStatus: session.StreamStatus,
	})
}

func streamStatusForEvent(event *livekitproto.WebhookEvent) (string, bool) {
	switch event.GetEvent() {
	case "participant_joined":
		return participantStreamStatus(event, "waiting_for_input")
	case "ingress_started":
		return "waiting_for_input", true
	case "track_published":
		return "live", true
	case "participant_left":
		return participantStreamStatus(event, "ending")
	case "ingress_ended":
		return "ending", true
	default:
		return "", false
	}
}

func participantStreamStatus(event *livekitproto.WebhookEvent, streamStatus string) (string, bool) {
	participant := event.GetParticipant()
	if participant == nil {
		return "", false
	}

	if participant.GetIsPublisher() || participant.GetPermission().GetCanPublish() {
		return streamStatus, true
	}

	return "", false
}

func webhookRoomName(event *livekitproto.WebhookEvent) string {
	if event.GetRoom().GetName() != "" {
		return event.GetRoom().GetName()
	}

	if event.GetIngressInfo().GetRoomName() != "" {
		return event.GetIngressInfo().GetRoomName()
	}

	return ""
}

func errIsWebhookInvalid(err error) bool {
	return errors.Is(err, livekit.ErrInvalidWebhook)
}
