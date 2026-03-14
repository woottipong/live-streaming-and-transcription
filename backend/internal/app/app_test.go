package app

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateSession(t *testing.T) {
	repo := &stubSessionRepository{
		createFn: func(_ context.Context, params CreateSessionParams) (Session, error) {
			if params.RoomName == "" {
				t.Fatalf("expected generated room name")
			}

			now := time.Date(2026, 3, 15, 1, 2, 3, 0, time.UTC)
			return Session{
				ID:                  uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				Title:               params.Title,
				Description:         params.Description,
				RoomName:            params.RoomName,
				SourceType:          params.SourceType,
				Language:            params.Language,
				ASRProvider:         params.ASRProvider,
				SubtitleEnabled:     params.SubtitleEnabled,
				StreamStatus:        "created",
				TranscriptionStatus: "idle",
				CreatedAt:           now,
				UpdatedAt:           now,
			}, nil
		},
	}

	app := New(repo)

	body := bytes.NewBufferString(`{"title":"Town Hall","language":"th-TH","source_type":"browser","asr_provider":"deepgram","subtitle_enabled":true}`)
	req := httptest.NewRequest(http.MethodPost, "/api/sessions", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test returned error: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode)
	}

	var got Session
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got.Title != "Town Hall" || got.Language != "th-TH" || !got.SubtitleEnabled {
		t.Fatalf("unexpected response body: %+v", got)
	}
	if got.StreamStatus != "created" || got.TranscriptionStatus != "idle" {
		t.Fatalf("expected default statuses, got %+v", got)
	}
}

func TestCreateSessionValidation(t *testing.T) {
	app := New(&stubSessionRepository{})

	req := httptest.NewRequest(http.MethodPost, "/api/sessions", strings.NewReader(`{"language":"th-TH"}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test returned error: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", resp.StatusCode)
	}
}

func TestGetSessionByID(t *testing.T) {
	sessionID := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	now := time.Date(2026, 3, 15, 1, 2, 3, 0, time.UTC)
	repo := &stubSessionRepository{
		getByIDFn: func(_ context.Context, id uuid.UUID) (Session, error) {
			if id != sessionID {
				t.Fatalf("unexpected ID: %s", id)
			}

			return Session{
				ID:                  sessionID,
				Title:               "Town Hall",
				Description:         stringPtr("Weekly sync"),
				RoomName:            "session-town-hall",
				SourceType:          "browser",
				Language:            "th-TH",
				ASRProvider:         "deepgram",
				SubtitleEnabled:     true,
				StreamStatus:        "created",
				TranscriptionStatus: "idle",
				CreatedAt:           now,
				UpdatedAt:           now,
			}, nil
		},
	}

	app := New(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/sessions/"+sessionID.String(), nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var got Session
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got.ID != sessionID || got.Description == nil || *got.Description != "Weekly sync" {
		t.Fatalf("unexpected response body: %+v", got)
	}
}

func TestGetSessionNotFound(t *testing.T) {
	repo := &stubSessionRepository{
		getByIDFn: func(_ context.Context, _ uuid.UUID) (Session, error) {
			return Session{}, ErrSessionNotFound
		},
	}

	app := New(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/sessions/33333333-3333-3333-3333-333333333333", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.StatusCode)
	}
}

func TestPatchSession(t *testing.T) {
	sessionID := uuid.MustParse("44444444-4444-4444-4444-444444444444")
	now := time.Date(2026, 3, 15, 1, 2, 3, 0, time.UTC)
	repo := &stubSessionRepository{
		updateFn: func(_ context.Context, id uuid.UUID, params UpdateSessionParams) (Session, error) {
			if id != sessionID {
				t.Fatalf("unexpected ID: %s", id)
			}
			if params.Title == nil || *params.Title != "Updated title" {
				t.Fatalf("expected title patch, got %+v", params)
			}

			return Session{
				ID:                  sessionID,
				Title:               *params.Title,
				Description:         params.Description,
				RoomName:            "session-updated-title",
				SourceType:          "browser",
				Language:            "th-TH",
				ASRProvider:         "deepgram",
				SubtitleEnabled:     true,
				StreamStatus:        "created",
				TranscriptionStatus: "idle",
				CreatedAt:           now,
				UpdatedAt:           now.Add(time.Minute),
			}, nil
		},
	}

	app := New(repo)
	req := httptest.NewRequest(http.MethodPatch, "/api/sessions/"+sessionID.String(), strings.NewReader(`{"title":"Updated title","description":"Updated description"}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var got Session
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got.Title != "Updated title" || got.Description == nil || *got.Description != "Updated description" {
		t.Fatalf("unexpected response body: %+v", got)
	}
}

func TestPatchSessionValidation(t *testing.T) {
	app := New(&stubSessionRepository{})

	req := httptest.NewRequest(http.MethodPatch, "/api/sessions/44444444-4444-4444-4444-444444444444", strings.NewReader(`{"language":"   "}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test returned error: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", resp.StatusCode)
	}
}

func TestPatchSessionClearsDescription(t *testing.T) {
	sessionID := uuid.MustParse("55555555-5555-5555-5555-555555555555")
	now := time.Date(2026, 3, 15, 1, 2, 3, 0, time.UTC)
	repo := &stubSessionRepository{
		updateFn: func(_ context.Context, id uuid.UUID, params UpdateSessionParams) (Session, error) {
			if id != sessionID {
				t.Fatalf("unexpected ID: %s", id)
			}
			if !params.ClearDescription {
				t.Fatalf("expected description to be cleared, got %+v", params)
			}
			if params.Description != nil {
				t.Fatalf("expected nil description when clearing, got %+v", params)
			}

			return Session{
				ID:                  sessionID,
				Title:               "Town Hall",
				Description:         nil,
				RoomName:            "session-town-hall",
				SourceType:          "browser",
				Language:            "th-TH",
				ASRProvider:         "deepgram",
				SubtitleEnabled:     true,
				StreamStatus:        "created",
				TranscriptionStatus: "idle",
				CreatedAt:           now,
				UpdatedAt:           now.Add(time.Minute),
			}, nil
		},
	}

	app := New(repo)
	req := httptest.NewRequest(http.MethodPatch, "/api/sessions/"+sessionID.String(), strings.NewReader(`{"description":null}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var got Session
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got.Description != nil {
		t.Fatalf("expected cleared description, got %+v", got)
	}
}

type stubSessionRepository struct {
	createFn  func(context.Context, CreateSessionParams) (Session, error)
	getByIDFn func(context.Context, uuid.UUID) (Session, error)
	updateFn  func(context.Context, uuid.UUID, UpdateSessionParams) (Session, error)
}

func (s *stubSessionRepository) Create(ctx context.Context, params CreateSessionParams) (Session, error) {
	if s.createFn == nil {
		return Session{}, errors.New("unexpected create call")
	}

	return s.createFn(ctx, params)
}

func (s *stubSessionRepository) GetByID(ctx context.Context, id uuid.UUID) (Session, error) {
	if s.getByIDFn == nil {
		return Session{}, errors.New("unexpected get call")
	}

	return s.getByIDFn(ctx, id)
}

func (s *stubSessionRepository) Update(ctx context.Context, id uuid.UUID, params UpdateSessionParams) (Session, error) {
	if s.updateFn == nil {
		return Session{}, errors.New("unexpected update call")
	}

	return s.updateFn(ctx, id, params)
}

func stringPtr(value string) *string {
	return &value
}
