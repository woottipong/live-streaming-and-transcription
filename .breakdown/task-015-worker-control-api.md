# Task: Worker Control API

## Task ID
task-015

## Epic
epic-07-transcription-worker

## Area
worker

## Status
todo

## Priority
high

## Depends On
- task-014

## Summary
เชื่อม /start และ /stop ให้ orchestrate full transcription pipeline (join room → subscribe audio → connect ASR → stream)

## Scope
- POST /start → สร้าง transcription session: join room + subscribe audio + connect ASR + เริ่ม stream
- POST /stop → หยุด transcription: flush + disconnect ASR + leave room + cleanup
- session lifecycle management (ป้องกัน double start, handle concurrent requests)
- ส่ง status update กลับ backend (POST /internal/transcripts/event ด้วย type "status")
- error handling: ถ้า step ใด fail ให้ cleanup step ก่อนหน้า

## Out of Scope
- auto-restart
- multi-worker coordination
- health monitoring

## Acceptance Criteria
- POST /start → full pipeline ทำงาน (join + subscribe + ASR + stream)
- POST /stop → cleanup ทุกอย่าง
- double start ไม่ทำให้ crash (return error หรือ idempotent)
- partial failure → cleanup อย่างถูกต้อง

## Files Likely Affected
- worker/internal/session/manager.go
- worker/internal/handler/

## Test Checklist
- [ ] POST /start → full pipeline ทำงาน
- [ ] POST /stop → หยุดทุกอย่าง clean
- [ ] double POST /start → ไม่ crash
- [ ] kill publisher ระหว่าง transcription → worker handle ได้

## Outcome


## Completion Evidence


## Completed At

