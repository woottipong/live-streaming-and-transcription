# Task: Database Schema

## Task ID
task-002

## Epic
epic-02-session-management

## Area
backend

## Status
todo

## Priority
high

## Depends On
- task-001

## Summary
สร้าง database schema สำหรับ sessions และ transcript_segments ตาม data model ใน ARCHITECTURE.md

## Scope
- สร้าง migration files
- สร้างตาราง `sessions` ครบทุก field (id, title, description, room_name, ingress_id, source_type, language, asr_provider, subtitle_enabled, stream_status, transcription_status, started_at, ended_at, created_at, updated_at)
- สร้างตาราง `transcript_segments` ครบทุก field (id, session_id, seq, text, start_time_ms, end_time_ms, speaker_id, language, confidence, provider, created_at)
- default values: stream_status = 'created', transcription_status = 'idle'
- indexes ที่จำเป็น

## Out of Scope
- ORM model
- repository layer

## Acceptance Criteria
- migration run ได้สำเร็จ
- ตาราง sessions มีครบทุก field
- ตาราง transcript_segments มีครบทุก field
- foreign key session_id อ้าง sessions.id
- default values ทำงานถูกต้อง

## Files Likely Affected
- backend/migrations/

## Test Checklist
- [ ] migrate up สำเร็จ
- [ ] migrate down สำเร็จ
- [ ] INSERT session ได้โดยไม่ต้องระบุ stream_status (ได้ default 'created')
- [ ] INSERT transcript_segment ได้โดยอ้าง session_id ที่มีอยู่

## Outcome


## Completion Evidence


## Completed At

