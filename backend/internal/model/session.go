package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID                  uuid.UUID  `json:"id"`
	Title               string     `json:"title"`
	Description         *string    `json:"description,omitempty"`
	RoomName            string     `json:"room_name"`
	IngressID           *string    `json:"ingress_id,omitempty"`
	SourceType          string     `json:"source_type"`
	Language            string     `json:"language"`
	ASRProvider         string     `json:"asr_provider"`
	SubtitleEnabled     bool       `json:"subtitle_enabled"`
	StreamStatus        string     `json:"stream_status"`
	TranscriptionStatus string     `json:"transcription_status"`
	StartedAt           *time.Time `json:"started_at,omitempty"`
	EndedAt             *time.Time `json:"ended_at,omitempty"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}
