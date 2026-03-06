# Task: LiveKit Room Creation

## Task ID
task-004

## Epic
epic-03-livekit-integration

## Area
backend

## Status
todo

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
- [ ] สร้าง session แล้ว room ปรากฏใน LiveKit
- [ ] room_name ถูก save ลง DB
- [ ] ถ้า LiveKit down ได้ error response ที่เข้าใจได้

## Outcome


## Completion Evidence


## Completed At

