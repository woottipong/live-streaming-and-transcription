# Task: Transcript Persist

## Task ID
task-019

## Epic
epic-08-transcript-delivery

## Area
backend

## Status
todo

## Priority
high

## Depends On
- task-002
- task-018

## Summary
Persist final transcript segments ลง PostgreSQL

## Scope
- save final transcript event เป็น row ใน transcript_segments table
- map event fields → DB columns (sessionId→session_id, seq, text, startTimeMs→start_time_ms, etc.)
- handle duplicate seq (upsert หรือ ignore)
- repository layer สำหรับ transcript_segments
- GET /api/sessions/:id/transcripts — ดึง transcript ของ session (ordered by seq)

## Out of Scope
- pagination
- full-text search
- export

## Acceptance Criteria
- final event ถูก save ลง DB ถูกต้อง
- ดึง transcript ของ session ได้เรียงตาม seq
- duplicate seq ไม่ทำให้ error
- field mapping ถูกต้องครบ

## Files Likely Affected
- backend/internal/repository/transcript.go
- backend/internal/handler/transcript.go

## Test Checklist
- [ ] final event → row ปรากฏใน DB
- [ ] fields ถูกต้องครบ
- [ ] GET /api/sessions/:id/transcripts → ได้ transcript เรียงตาม seq
- [ ] duplicate seq → ไม่ error

## Outcome


## Completion Evidence


## Completed At

