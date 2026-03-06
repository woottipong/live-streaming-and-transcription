# Task: Auto Start Transcription

## Task ID
task-022

## Epic
epic-03-livekit-integration

## Area
backend

## Status
todo

## Priority
high

## Depends On
- task-006
- task-015

## Summary
Backend auto-trigger transcription worker เมื่อได้รับ LiveKit webhook TrackPublished (audio)

## Scope
- เมื่อ webhook TrackPublished มาถึง (audio track):
  1. ตรวจว่า session มี subtitle_enabled = true
  2. ตรวจว่า transcription_status = idle (ยังไม่เริ่ม)
  3. POST /start ไปหา Worker (WORKER_URL) พร้อม sessionId, roomName, asrProvider, language
  4. update transcription_status เป็น 'starting'
- เมื่อ webhook ParticipantLeft / IngressEnded มาถึง:
  1. ตรวจว่า transcription_status = running
  2. POST /stop ไปหา Worker
  3. update transcription_status เป็น 'stopping'
- config: WORKER_URL จาก env
- handle Worker unreachable (log error, set transcription_status = failed)

## Out of Scope
- retry logic
- worker health check
- manual start/stop API

## Acceptance Criteria
- publisher publish audio → transcription เริ่มอัตโนมัติ (ถ้า subtitle_enabled)
- publisher leave → transcription หยุดอัตโนมัติ
- subtitle_enabled = false → ไม่ trigger worker
- worker unreachable → transcription_status = failed
- transcription ที่กำลังทำอยู่ → ไม่ trigger ซ้ำ

## Files Likely Affected
- backend/internal/handler/webhook.go
- backend/internal/service/transcription.go

## Test Checklist
- [ ] publish audio + subtitle_enabled → worker ได้รับ /start
- [ ] publish audio + subtitle_enabled=false → ไม่เกิดอะไร
- [ ] publisher leave → worker ได้รับ /stop
- [ ] worker down → transcription_status = failed
- [ ] double TrackPublished → ไม่ trigger ซ้ำ

## Outcome


## Completion Evidence


## Completed At

