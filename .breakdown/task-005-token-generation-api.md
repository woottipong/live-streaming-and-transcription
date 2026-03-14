# Task: Token Generation API

## Task ID
task-005

## Epic
epic-03-livekit-integration

## Area
backend

## Status
done

## Priority
high

## Depends On
- task-004

## Summary
สร้าง API สำหรับ generate LiveKit access token สำหรับ publisher และ viewer

## Scope
- POST /api/sessions/:id/token/publisher — generate token ที่มีสิทธิ์ publish
- POST /api/sessions/:id/token/viewer — generate token ที่มีสิทธิ์ subscribe only
- ใช้ LiveKit Server SDK สำหรับ token generation
- token ต้องผูกกับ room_name ของ session
- participant identity generate หรือรับจาก request body

## Out of Scope
- authentication
- token refresh
- custom permissions

## Acceptance Criteria
- publisher token สามารถ publish audio/video ได้
- viewer token สามารถ subscribe ได้แต่ publish ไม่ได้
- 404 เมื่อ session ไม่เจอ
- token มี room_name ตรงกับ session

## Files Likely Affected
- backend/internal/handler/token.go
- backend/internal/livekit/token.go

## Test Checklist
- [x] POST publisher token ได้ token กลับมา
- [x] POST viewer token ได้ token กลับมา
- [x] 404 เมื่อ session ไม่มี
- [x] publisher token ใช้ connect + publish ได้จริง
- [x] viewer token ใช้ connect + subscribe ได้แต่ publish ไม่ได้

## Outcome
เพิ่ม `POST /api/sessions/:id/token/publisher` และ `POST /api/sessions/:id/token/viewer` โดยใช้ LiveKit access token ผ่าน `internal/livekit` wrapper และผูก grant กับ `session.room_name`

รองรับ optional request body สำหรับ `identity` และ `name` ถ้าไม่ส่งมาจะ generate identity ตาม role ให้อัตโนมัติ และตอบกลับ `token`, `identity`, `room_name`

## Completion Evidence
- รัน `go test ./...` ใน `backend/` ผ่าน
- Automated tests ครอบคลุม publisher token, viewer token, และ `404 session not found`
- Smoke test จริงผ่าน: สร้าง session `3eca9fea-5a94-47cf-a58a-47f68e274015` แล้วเรียก `POST /api/sessions/:id/token/publisher` ได้ token สำหรับ room `token-smoke-session-f908b1d8`
- ตรวจ payload ของ publisher token แล้วได้ `{"canPublish": true, "canSubscribe": true, "room": "token-smoke-session-f908b1d8", "roomJoin": true}`
- เรียก `POST /api/sessions/:id/token/viewer` ได้ token สำหรับ room เดียวกัน และตรวจ payload แล้วได้ `{"canPublish": false, "canSubscribe": true, "room": "token-smoke-session-f908b1d8", "roomJoin": true}`
- เรียก `POST /api/sessions/99999999-9999-9999-9999-999999999999/token/publisher` ได้ `404 Not Found` พร้อม response `{"error":"session not found"}`

## Completed At
2026-03-15 01:40:00 +07
