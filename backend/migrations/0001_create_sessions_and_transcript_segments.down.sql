DROP TRIGGER IF EXISTS trigger_sessions_set_updated_at ON sessions;
DROP FUNCTION IF EXISTS set_updated_at_timestamp();
DROP TABLE IF EXISTS transcript_segments;
DROP TABLE IF EXISTS sessions;
