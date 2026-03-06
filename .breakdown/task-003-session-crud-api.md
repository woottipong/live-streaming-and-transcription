# Task: Session CRUD API

## Task ID
task-003

## Epic
epic-02-session-management

## Area
backend

## Status
todo

## Priority
high

## Depends On
- task-002

## Summary
ทำ Session CRUD API (POST, GET, PATCH) ตาม API design ใน ARCHITECTURE.md

## Scope
- POST /api/sessions — สร้าง session ใหม่
- GET /api/sessions/:id — ดึง session by ID
- PATCH /api/sessions/:id — update session fields
- database connection setup (pgx หรือ sqlx)
- request validation
- proper error responses (400, 404, 500)
- response format เป็น JSON

## Out of Scope
- authentication
- pagination
- list endpoint
- LiveKit room creation

## Acceptance Criteria
- POST สร้าง session ได้ ส่ง title, language, source_type, asr_provider, subtitle_enabled
- GET ดึง session กลับมาได้ครบทุก field
- PATCH update ได้ (เช่น title, description)
- 404 เมื่อ session ไม่เจอ
- 400 เมื่อ input ไม่ถูกต้อง

## Files Likely Affected
- backend/internal/handler/
- backend/internal/repository/
- backend/internal/model/

## Test Checklist
- [ ] POST /api/sessions สร้าง session ได้
- [ ] GET /api/sessions/:id ดึงได้
- [ ] PATCH /api/sessions/:id update ได้
- [ ] 404 เมื่อ GET session ที่ไม่มี
- [ ] 400 เมื่อ POST ไม่มี required field

## Outcome


## Completion Evidence


## Completed At

