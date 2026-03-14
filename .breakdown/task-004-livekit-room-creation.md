# Task: LiveKit Room Creation

## Task ID
task-004

## Epic
epic-03-livekit-integration

## Area
backend

## Status
done

## Priority
high

## Depends On
- task-003

## Summary
เชื่อม LiveKit Server SDK เพื่อสร้าง room เมื่อสร้าง session

## Scope
- integrate LiveKit Server SDK (Go)
- สร้าง LiveKit room เมื่อ POST /api/sessions
- ใช้ room_name จาก session (generate หรือ derive จาก session ID)
- save room_name ลง session record
- config LiveKit host/api key/api secret จาก env

## Out of Scope
- ingress
- egress
- room deletion

## Acceptance Criteria
- POST /api/sessions สร้าง LiveKit room ได้จริง
- room_name ถูก save ลง DB
- response มี room_name
- ถ้า LiveKit ไม่พร้อม ให้ return error ชัดเจน

## Files Likely Affected
- backend/internal/livekit/
- backend/internal/handler/session.go

## Test Checklist
- [x] สร้าง session แล้ว room ปรากฏใน LiveKit
- [x] room_name ถูก save ลง DB
- [x] ถ้า LiveKit down ได้ error response ที่เข้าใจได้

## Outcome
เพิ่ม LiveKit room creation เข้า `POST /api/sessions` ผ่าน `internal/livekit` client โดยใช้ `LIVEKIT_HTTP_URL`, `LIVEKIT_API_KEY`, `LIVEKIT_API_SECRET` จาก env และทำ best-effort cleanup ของ room ถ้า save session ลง DB ไม่สำเร็จ

ปรับ backend wiring ให้ inject LiveKit client ผ่าน app layer และเพิ่ม automated tests ครอบคลุม success path, LiveKit unavailable (`503`), และ cleanup behavior

## Completion Evidence
- รัน `go test ./...` ใน `backend/` ผ่าน
- Smoke test จริงผ่าน: `POST /api/sessions` ตอบกลับ `room_name` และ query DB พบ `sessions.room_name = smoke-test-session-8b34110c` สำหรับ session `4c5e12e6-c738-4b25-a118-a237faae20cc`
- ตรวจผ่าน LiveKit RoomService API แล้วพบ room `smoke-test-session-8b34110c` ถูกสร้างจริง
- เมื่อ stop LiveKit ชั่วคราวแล้วเรียก `POST /api/sessions` ได้ `503 Service Unavailable` พร้อม response `{"error":"failed to create LiveKit room"}`

## Completed At
2026-03-15 00:47:00 +07
