# Task: Database Schema

## Task ID
task-002

## Epic
epic-02-session-management

## Area
backend

## Status
done

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
- [x] migrate up สำเร็จ
- [x] migrate down สำเร็จ
- [x] INSERT session ได้โดยไม่ต้องระบุ stream_status (ได้ default 'created')
- [x] INSERT transcript_segment ได้โดยอ้าง session_id ที่มีอยู่

## Outcome
เพิ่ม migration สำหรับ `sessions` และ `transcript_segments` พร้อม default values, constraints, foreign key, และ indexes ที่จำเป็น

เลือกใช้ `UUID` เป็น primary key สำหรับทั้งสองตาราง และเปิด `pgcrypto` เพื่อใช้ `gen_random_uuid()` ใน PostgreSQL


## Completion Evidence
รัน `0001_create_sessions_and_transcript_segments.up.sql` ผ่านใน PostgreSQL จริง
รัน `0001_create_sessions_and_transcript_segments.down.sql` ผ่าน และรัน `up` ซ้ำได้สำเร็จ
`INSERT INTO sessions (...) RETURNING id, stream_status, transcription_status` ได้ค่า default `created` และ `idle`
`INSERT INTO transcript_segments (...)` ผ่านเมื่ออ้าง `session_id` ของ `sessions` ที่มีอยู่
ตรวจ `information_schema.columns` แล้วยืนยัน field ครบทั้งสองตาราง และ foreign key `transcript_segments.session_id -> sessions.id` ทำงานถูกต้อง


## Completed At
2026-03-15 00:12:00 +07
