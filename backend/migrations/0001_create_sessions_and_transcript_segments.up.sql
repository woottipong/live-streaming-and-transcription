CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT,
    room_name TEXT NOT NULL UNIQUE,
    ingress_id TEXT,
    source_type TEXT NOT NULL CHECK (source_type IN ('browser', 'rtmp')),
    language VARCHAR(16) NOT NULL,
    asr_provider TEXT NOT NULL,
    subtitle_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    stream_status TEXT NOT NULL DEFAULT 'created' CHECK (stream_status IN ('created', 'ready', 'waiting_for_input', 'live', 'ending', 'ended', 'failed')),
    transcription_status TEXT NOT NULL DEFAULT 'idle' CHECK (transcription_status IN ('idle', 'starting', 'running', 'stopping', 'stopped', 'failed')),
    started_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT sessions_ended_after_started CHECK (ended_at IS NULL OR started_at IS NULL OR ended_at >= started_at)
);

CREATE INDEX idx_sessions_stream_status ON sessions (stream_status);
CREATE INDEX idx_sessions_transcription_status ON sessions (transcription_status);
CREATE INDEX idx_sessions_created_at ON sessions (created_at DESC);

CREATE FUNCTION set_updated_at_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_sessions_set_updated_at
    BEFORE UPDATE ON sessions
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at_timestamp();

CREATE TABLE transcript_segments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID NOT NULL REFERENCES sessions (id) ON DELETE CASCADE,
    seq BIGINT NOT NULL,
    text TEXT NOT NULL,
    start_time_ms BIGINT NOT NULL CHECK (start_time_ms >= 0),
    end_time_ms BIGINT NOT NULL CHECK (end_time_ms >= start_time_ms),
    speaker_id TEXT,
    language VARCHAR(16) NOT NULL,
    confidence NUMERIC(5,4) CHECK (confidence IS NULL OR (confidence >= 0 AND confidence <= 1)),
    provider TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT transcript_segments_session_seq_unique UNIQUE (session_id, seq)
);

CREATE INDEX idx_transcript_segments_session_id ON transcript_segments (session_id);
CREATE INDEX idx_transcript_segments_session_id_created_at ON transcript_segments (session_id, created_at);
