# Task: Transcript Event Receiver

## Task ID
task-018

## Epic
epic-08-transcript-delivery

## Area
backend

## Status
todo

## Priority
high

## Depends On
- task-017

## Summary
สร้าง internal API endpoint รับ transcript event จาก Worker แล้ว broadcast ผ่าน WS Hub

## Scope
- POST /internal/transcripts/event
- validate transcript event schema (sessionId, seq, type, text, startTimeMs, endTimeMs, language, confidence)
- ส่ง event เข้า WS Hub เพื่อ broadcast ไป viewers
- แยก logic: partial → broadcast only, final → broadcast + persist

## Out of Scope
- authentication ของ internal API
- rate limiting
- deduplication

## Acceptance Criteria
- POST event → viewers ได้รับผ่าน WebSocket
- partial event → broadcast only (ไม่ลง DB)
- final event → broadcast + persist
- invalid event → 400 error

## Files Likely Affected
- backend/internal/handler/transcript_event.go

## Test Checklist
- [ ] POST partial event → viewer ได้รับผ่าน WS
- [ ] POST final event → viewer ได้รับ + ลง DB
- [ ] POST invalid event → 400
- [ ] ไม่มี viewer → ไม่ error

## Outcome


## Completion Evidence


## Completed At

