# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a **Live Streaming + Realtime Transcription** platform (MVP). The system streams audio/video via LiveKit and delivers realtime transcriptions to viewers via WebSocket.

Stack:
- **Frontend:** Next.js + LiveKit React SDK
- **Backend:** Go + Fiber (REST API + WebSocket)
- **Transcription Worker:** Go (long-running HTTP service)
- **Media:** LiveKit (cloud)
- **Database:** PostgreSQL

## Architecture

See [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) for the full architecture. Key points:

### Component Communication
- **Backend → Worker:** HTTP POST to `WORKER_URL` (default `http://worker:8081`)
  - `POST /start` — start transcription `{ sessionId, roomName, asrProvider, language }`
  - `POST /stop`  — stop transcription and flush `{ sessionId }`
- **Worker → Backend:** HTTP POST `/internal/transcripts/event` per transcript event
- **Backend → Viewers:** WebSocket broadcast via WS Hub at `/ws/transcripts/:sessionId`

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
- `partial` events are broadcast only; `final` events are broadcast AND persisted to DB.

### Session State
Sessions have two independent status fields:
- `stream_status`: `created` → `ready` → `waiting_for_input` → `live` → `ending` → `ended` / `failed`
- `transcription_status`: `idle` → `starting` → `running` → `stopping` → `stopped` / `failed`

### Key Trigger: Start Transcription
Triggered by LiveKit webhook `TrackPublished` (audio track), NOT `ParticipantJoined`.

## Development Workflow

This project uses a task-breakdown system defined in [docs/WORKFLOW.md](docs/WORKFLOW.md):

- `docs/ARCHITECTURE.md` — source of truth for system design
- `.breakdown/STATUS.md` — source of truth for task progress
- `.breakdown/task-*.md` — individual task specs (1 task = 1 deliverable)

### Working on a Task
Always provide these three files as context when implementing:
1. `docs/ARCHITECTURE.md`
2. `.breakdown/STATUS.md`
3. The specific `task-*.md` being worked on

Stay strictly within the task's defined scope. When done, update both the task file (Status → done, add Outcome + Completion Evidence) and `.breakdown/STATUS.md`.

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/sessions` | Create session |
| GET | `/api/sessions/:id` | Get session |
| PATCH | `/api/sessions/:id` | Update session |
| POST | `/api/sessions/:id/token/publisher` | Get publisher token |
| POST | `/api/sessions/:id/token/viewer` | Get viewer token |
| POST | `/internal/livekit/webhook` | LiveKit webhook receiver |
| POST | `/internal/transcripts/event` | Receive transcript from worker |
| GET | `/ws/transcripts/:sessionId` | WebSocket transcript stream |

## Out of Scope (MVP)
Queue/background jobs, multi-provider ASR fallback, subtitle export (SRT/VTT), AI summary/translation, Redis, advanced analytics, speaker diarization.
