# AGENTS.md

This file provides guidance to coding agents working in this repository.

## Project Overview

This is a **Live Streaming + Realtime Transcription** platform (MVP). The system streams audio/video via LiveKit and delivers realtime transcriptions to viewers via WebSocket.

Stack:
- **Frontend:** Next.js + LiveKit React SDK
- **Backend:** Go + Fiber (REST API + WebSocket)
- **Transcription Worker:** Go (long-running HTTP service)
- **Media:** LiveKit
- **Database:** PostgreSQL

## Architecture

Use [docs/ARCHITECTURE.md](/Users/iwoody/Workspace/labs/realtime_streaming/docs/ARCHITECTURE.md) as the source of truth for system design and MVP boundaries.

Key points:

### Component Communication
- **Backend -> Worker:** HTTP POST to `WORKER_URL` (default `http://worker:8081`)
  - `POST /start` - start transcription `{ sessionId, roomName, asrProvider, language }`
  - `POST /stop` - stop transcription and flush `{ sessionId }`
- **Worker -> Backend:** HTTP POST `/internal/transcripts/event` per transcript event
- **Backend -> Viewers:** WebSocket broadcast via WS hub at `/ws/transcripts/:sessionId`

### Transcript Event Schema

```json
{
  "sessionId": "sess_123",
  "seq": 101,
  "type": "partial",
  "text": "...",
  "startTimeMs": 1000,
  "endTimeMs": 2400,
  "speakerId": "speaker_1",
  "language": "th-TH",
  "confidence": 0.89,
  "isFinal": false
}
```

- `partial` events are broadcast only
- `final` events are broadcast and persisted to DB

### Session State
- `stream_status`: `created` -> `ready` -> `waiting_for_input` -> `live` -> `ending` -> `ended` / `failed`
- `transcription_status`: `idle` -> `starting` -> `running` -> `stopping` -> `stopped` / `failed`

### Key Trigger
- Start transcription from LiveKit webhook `TrackPublished` for an audio track, not `ParticipantJoined`

## Frontend Commands

All frontend commands (lint, test, build, etc.) **must be run inside the Docker container**. Do not run `pnpm` or `node` directly on the host.

```bash
# Lint
docker compose exec frontend pnpm lint

# Build
docker compose exec frontend pnpm build

# Test (once a test script is added)
docker compose exec frontend pnpm test
```

If the container is not running, start it first:

```bash
docker compose up -d frontend
```

## Working Rules

- Work from these source-of-truth files:
  - [docs/ARCHITECTURE.md](/Users/iwoody/Workspace/labs/realtime_streaming/docs/ARCHITECTURE.md)
  - [.breakdown/STATUS.md](/Users/iwoody/Workspace/labs/realtime_streaming/.breakdown/STATUS.md)
  - the active [.breakdown/task-*.md](/Users/iwoody/Workspace/labs/realtime_streaming/.breakdown)
- Stay inside the active task scope. Do not expand beyond MVP.
- When a task is completed, update both the task file and [.breakdown/STATUS.md](/Users/iwoody/Workspace/labs/realtime_streaming/.breakdown/STATUS.md).
- Keep changes tightly scoped. Do not refactor unrelated code.
- Verify changes before claiming completion.

## API Endpoints

| Method | Path | Description |
|------|------|-------------|
| POST | `/api/sessions` | Create session |
| GET | `/api/sessions/:id` | Get session |
| PATCH | `/api/sessions/:id` | Update session |
| POST | `/api/sessions/:id/token/publisher` | Get publisher token |
| POST | `/api/sessions/:id/token/viewer` | Get viewer token |
| POST | `/internal/livekit/webhook` | LiveKit webhook receiver |
| POST | `/internal/transcripts/event` | Receive transcript from worker |
| GET | `/ws/transcripts/:sessionId` | WebSocket transcript stream |

## Out of Scope

Queue/background jobs, multi-provider ASR fallback, subtitle export (SRT/VTT), AI summary/translation, Redis, advanced analytics, and speaker diarization.
