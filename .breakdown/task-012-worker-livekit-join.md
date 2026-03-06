# Task: Worker LiveKit Join

## Task ID
task-012

## Epic
epic-07-transcription-worker

## Area
worker

## Status
todo

## Priority
high

## Depends On
- task-011
- task-004

## Summary
Worker join LiveKit room เป็น bot participant เมื่อได้รับ /start command

## Scope
- ใช้ LiveKit Server SDK สร้าง bot token (ไม่มีสิทธิ์ publish, มีสิทธิ์ subscribe)
- connect เข้า LiveKit room ด้วย roomName จาก /start payload
- participant identity เป็น "transcription-bot-{sessionId}"
- เก็บ room connection state ต่อ session (map[sessionId]*room)
- /stop → disconnect จาก room
- handle reconnect กรณี connection drop (basic)

## Out of Scope
- audio processing
- ASR connection
- multi-room per worker

## Acceptance Criteria
- POST /start → worker ปรากฏเป็น participant ใน LiveKit room
- participant identity ถูกต้อง
- POST /stop → worker ออกจาก room
- worker handle room ที่ไม่มีอยู่ได้ (return error)

## Files Likely Affected
- worker/internal/session/
- worker/internal/livekit/

## Test Checklist
- [ ] POST /start → เห็น bot ใน LiveKit dashboard
- [ ] identity เป็น transcription-bot-{sessionId}
- [ ] POST /stop → bot หายจาก room
- [ ] start ด้วย room ที่ไม่มี → error response

## Outcome


## Completion Evidence


## Completed At

