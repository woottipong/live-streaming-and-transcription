# Task: Token Generation API

## Task ID
task-005

## Epic
epic-03-livekit-integration

## Area
backend

## Status
todo

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
- [ ] POST publisher token ได้ token กลับมา
- [ ] POST viewer token ได้ token กลับมา
- [ ] 404 เมื่อ session ไม่มี
- [ ] publisher token ใช้ connect + publish ได้จริง
- [ ] viewer token ใช้ connect + subscribe ได้แต่ publish ไม่ได้

## Outcome


## Completion Evidence


## Completed At

