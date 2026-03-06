# Task: Worker Scaffold

## Task ID
task-011

## Epic
epic-07-transcription-worker

## Area
worker

## Status
todo

## Priority
high

## Depends On
- task-001

## Summary
ตั้งโครงสร้าง Transcription Worker เป็น Go HTTP service ที่รับ command จาก Backend

## Scope
- Worker เป็น long-running HTTP service ฟัง port 8081
- POST /start endpoint (รับ sessionId, roomName, asrProvider, language) — ตอนนี้แค่ log + return 200
- POST /stop endpoint (รับ sessionId) — ตอนนี้แค่ log + return 200
- GET /health endpoint
- config จาก env: LIVEKIT_URL, LIVEKIT_API_KEY, LIVEKIT_API_SECRET, BACKEND_URL, PORT
- graceful shutdown

## Out of Scope
- LiveKit join
- ASR connection
- actual transcription logic

## Acceptance Criteria
- `go run ./worker` ฟัง port 8081
- POST /start return 200 + log session info
- POST /stop return 200 + log session info
- GET /health return 200
- graceful shutdown ทำงาน

## Files Likely Affected
- worker/main.go
- worker/internal/handler/

## Test Checklist
- [ ] worker start ได้ไม่ error
- [ ] POST /start return 200
- [ ] POST /stop return 200
- [ ] GET /health return 200
- [ ] SIGTERM → graceful shutdown

## Outcome


## Completion Evidence


## Completed At

