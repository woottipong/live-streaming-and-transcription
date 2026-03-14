package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/iwoody/realtime-streaming/backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrSessionNotFound    = errors.New("session not found")
	ErrDuplicateRoomName  = errors.New("room name already exists")
	ErrConstraintViolated = errors.New("constraint violated")
)

type CreateSessionParams struct {
	Title           string
	Description     *string
	RoomName        string
	SourceType      string
	Language        string
	ASRProvider     string
	SubtitleEnabled bool
}

type UpdateSessionParams struct {
	Title            *string
	Description      *string
	ClearDescription bool
	Language         *string
	SourceType       *string
	ASRProvider      *string
	SubtitleEnabled  *bool
}

type SessionRepository interface {
	Create(context.Context, CreateSessionParams) (model.Session, error)
	GetByID(context.Context, uuid.UUID) (model.Session, error)
	GetByRoomName(context.Context, string) (model.Session, error)
	Update(context.Context, uuid.UUID, UpdateSessionParams) (model.Session, error)
	UpdateStreamStatusByRoomName(context.Context, string, string) (model.Session, error)
}

type PostgresSessionRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresSessionRepository(pool *pgxpool.Pool) *PostgresSessionRepository {
	return &PostgresSessionRepository{pool: pool}
}

func NewPostgresPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return pool, nil
}

func (r *PostgresSessionRepository) Create(ctx context.Context, params CreateSessionParams) (model.Session, error) {
	const query = `
		INSERT INTO sessions (
			title,
			description,
			room_name,
			source_type,
			language,
			asr_provider,
			subtitle_enabled
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, title, description, room_name, ingress_id, source_type, language, asr_provider,
			subtitle_enabled, stream_status, transcription_status, started_at, ended_at, created_at, updated_at
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		params.Title,
		params.Description,
		params.RoomName,
		params.SourceType,
		params.Language,
		params.ASRProvider,
		params.SubtitleEnabled,
	)

	session, err := scanSession(row)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return model.Session{}, ErrDuplicateRoomName
		}
		return model.Session{}, fmt.Errorf("create session: %w", err)
	}

	return session, nil
}

func (r *PostgresSessionRepository) GetByID(ctx context.Context, id uuid.UUID) (model.Session, error) {
	const query = `
		SELECT id, title, description, room_name, ingress_id, source_type, language, asr_provider,
			subtitle_enabled, stream_status, transcription_status, started_at, ended_at, created_at, updated_at
		FROM sessions
		WHERE id = $1
	`

	session, err := scanSession(r.pool.QueryRow(ctx, query, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Session{}, ErrSessionNotFound
	}
	if err != nil {
		return model.Session{}, fmt.Errorf("get session by id: %w", err)
	}

	return session, nil
}

func (r *PostgresSessionRepository) GetByRoomName(ctx context.Context, roomName string) (model.Session, error) {
	const query = `
		SELECT id, title, description, room_name, ingress_id, source_type, language, asr_provider,
			subtitle_enabled, stream_status, transcription_status, started_at, ended_at, created_at, updated_at
		FROM sessions
		WHERE room_name = $1
	`

	session, err := scanSession(r.pool.QueryRow(ctx, query, roomName))
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Session{}, ErrSessionNotFound
	}
	if err != nil {
		return model.Session{}, fmt.Errorf("get session by room name: %w", err)
	}

	return session, nil
}

func (r *PostgresSessionRepository) Update(ctx context.Context, id uuid.UUID, params UpdateSessionParams) (model.Session, error) {
	const query = `
		UPDATE sessions
		SET
			title = COALESCE($2, title),
			description = CASE
				WHEN $3 THEN NULL
				WHEN $4 IS NOT NULL THEN $4
				ELSE description
			END,
			language = COALESCE($5, language),
			source_type = COALESCE($6, source_type),
			asr_provider = COALESCE($7, asr_provider),
			subtitle_enabled = COALESCE($8, subtitle_enabled)
		WHERE id = $1
		RETURNING id, title, description, room_name, ingress_id, source_type, language, asr_provider,
			subtitle_enabled, stream_status, transcription_status, started_at, ended_at, created_at, updated_at
	`

	session, err := scanSession(r.pool.QueryRow(
		ctx,
		query,
		id,
		params.Title,
		params.ClearDescription,
		params.Description,
		params.Language,
		params.SourceType,
		params.ASRProvider,
		params.SubtitleEnabled,
	))
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Session{}, ErrSessionNotFound
	}
	if err != nil {
		if strings.Contains(err.Error(), "violates check constraint") || strings.Contains(err.Error(), "duplicate key value") {
			return model.Session{}, fmt.Errorf("%w: %s", ErrConstraintViolated, err.Error())
		}
		return model.Session{}, fmt.Errorf("update session: %w", err)
	}

	return session, nil
}

func (r *PostgresSessionRepository) UpdateStreamStatusByRoomName(ctx context.Context, roomName string, streamStatus string) (model.Session, error) {
	const query = `
		UPDATE sessions
		SET stream_status = $2
		WHERE room_name = $1
		RETURNING id, title, description, room_name, ingress_id, source_type, language, asr_provider,
			subtitle_enabled, stream_status, transcription_status, started_at, ended_at, created_at, updated_at
	`

	session, err := scanSession(r.pool.QueryRow(ctx, query, roomName, streamStatus))
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Session{}, ErrSessionNotFound
	}
	if err != nil {
		if strings.Contains(err.Error(), "violates check constraint") {
			return model.Session{}, fmt.Errorf("%w: %s", ErrConstraintViolated, err.Error())
		}
		return model.Session{}, fmt.Errorf("update session stream status by room name: %w", err)
	}

	return session, nil
}

func scanSession(row pgx.Row) (model.Session, error) {
	var session model.Session

	err := row.Scan(
		&session.ID,
		&session.Title,
		&session.Description,
		&session.RoomName,
		&session.IngressID,
		&session.SourceType,
		&session.Language,
		&session.ASRProvider,
		&session.SubtitleEnabled,
		&session.StreamStatus,
		&session.TranscriptionStatus,
		&session.StartedAt,
		&session.EndedAt,
		&session.CreatedAt,
		&session.UpdatedAt,
	)
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}
