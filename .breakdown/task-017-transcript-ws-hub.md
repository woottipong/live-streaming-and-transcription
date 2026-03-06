# Task: Transcript WebSocket Hub

## Task ID
task-017

## Epic
epic-08-transcript-delivery

## Area
backend

## Status
todo

## Priority
high

## Depends On
- task-003

## Summary
สร้าง WebSocket Hub ใน Backend สำหรับ broadcast transcript events ไปยัง viewers

## Scope
- GET /ws/transcripts/:sessionId — WebSocket endpoint
- Hub เก็บ connections แยกตาม sessionId
- support multiple viewers per session
- broadcast message ไปทุก viewer ที่ subscribe session เดียวกัน
- handle client connect / disconnect / cleanup
- graceful close เมื่อ session end

## Out of Scope
- authentication
- message history/replay
- binary protocol
- Redis pub/sub

## Acceptance Criteria
- viewer connect WebSocket ได้
- multiple viewers เห็น message เดียวกัน
- viewer disconnect → connection ถูก cleanup
- broadcast ไม่ block ถ้า viewer ช้า (async send หรือ drop)

## Files Likely Affected
- backend/internal/ws/
- backend/internal/handler/transcript_ws.go

## Test Checklist
- [ ] connect WebSocket ได้
- [ ] 2 viewers เห็น broadcast เดียวกัน
- [ ] viewer disconnect → ไม่ leak connection
- [ ] send message เร็วๆ → ไม่ block

## Outcome


## Completion Evidence


## Completed At

