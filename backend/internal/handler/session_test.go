package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/iwoody/realtime-streaming/backend/internal/model"
	"github.com/iwoody/realtime-streaming/backend/internal/repository"
)

func TestCreateSessionWithLiveKitUsesDetachedContextForCleanup(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	roomClient := &stubRoomClient{
		createRoomFn: func(context.Context, string) error {
			return nil
		},
		deleteRoomFn: func(cleanupCtx context.Context, _ string) error {
			if cleanupCtx.Err() != nil {
				t.Fatalf("expected cleanup context to stay active, got %v", cleanupCtx.Err())
			}

			return nil
		},
	}
	handler := NewSessionHandler(&stubSessionRepository{
		createFn: func(context.Context, repository.CreateSessionParams) (model.Session, error) {
			cancel()
			return model.Session{}, context.Canceled
		},
	}, roomClient, roomClient)

	_, err := handler.createSessionWithLiveKit(ctx, repository.CreateSessionParams{
		Title:       "Town Hall",
		RoomName:    "town-hall-12345678",
		SourceType:  "browser",
		Language:    "th-TH",
		ASRProvider: "deepgram",
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled error, got %v", err)
	}
}

type stubRoomClient struct {
	createRoomFn func(context.Context, string) error
	deleteRoomFn func(context.Context, string) error
}

func (s *stubRoomClient) CreateRoom(ctx context.Context, roomName string) error {
	if s.createRoomFn == nil {
		return nil
	}

	return s.createRoomFn(ctx, roomName)
}

func (s *stubRoomClient) DeleteRoom(ctx context.Context, roomName string) error {
	if s.deleteRoomFn == nil {
		return nil
	}

	return s.deleteRoomFn(ctx, roomName)
}

func (*stubRoomClient) CreatePublisherToken(string, string, string) (string, error) {
	return "", nil
}

func (*stubRoomClient) CreateViewerToken(string, string, string) (string, error) {
	return "", nil
}

type stubSessionRepository struct {
	createFn func(context.Context, repository.CreateSessionParams) (model.Session, error)
}

func (s *stubSessionRepository) Create(ctx context.Context, params repository.CreateSessionParams) (model.Session, error) {
	if s.createFn == nil {
		return model.Session{}, errors.New("unexpected create call")
	}

	return s.createFn(ctx, params)
}

func (*stubSessionRepository) GetByID(context.Context, uuid.UUID) (model.Session, error) {
	return model.Session{}, errors.New("unexpected get by id call")
}

func (*stubSessionRepository) Update(context.Context, uuid.UUID, repository.UpdateSessionParams) (model.Session, error) {
	return model.Session{}, errors.New("unexpected update call")
}
