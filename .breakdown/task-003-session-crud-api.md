# Task: Session CRUD API

## Task ID
task-003

## Epic
epic-02-session-management

## Area
backend

## Status
done

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
- [x] POST /api/sessions สร้าง session ได้
- [x] GET /api/sessions/:id ดึงได้
- [x] PATCH /api/sessions/:id update ได้
- [x] 404 เมื่อ GET session ที่ไม่มี
- [x] 400 เมื่อ POST ไม่มี required field

## Outcome
เพิ่ม Session CRUD API ครบ 3 endpoint พร้อม request validation, JSON error responses, และ PostgreSQL repository ผ่าน `pgxpool`

สร้าง Fiber app wiring ใหม่ให้ backend bootstrap ด้วย database connection และแยกโค้ดเป็น `internal/app`, `internal/handler`, `internal/repository`, `internal/model`

## Completion Evidence
รัน `go test ./...` ใน `backend/` ผ่าน
ทดสอบ behavior ของ `POST /api/sessions`, `GET /api/sessions/:id`, `PATCH /api/sessions/:id` ผ่านด้วย handler tests
ทดสอบกรณี validation error (`400`) และไม่พบ session (`404`) ผ่าน
POST สร้าง `room_name` อัตโนมัติจาก backend โดยไม่บังคับให้ client ส่งมา

## Completed At
2026-03-15 01:05:00 +07
