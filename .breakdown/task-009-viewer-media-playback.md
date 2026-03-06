# Task: Viewer Media Playback

## Task ID
task-009

## Epic
epic-05-viewer-flow

## Area
frontend

## Status
todo

## Priority
high

## Depends On
- task-005

## Summary
ทำหน้า viewer ให้ connect LiveKit room แล้วดู live video/audio

## Scope
- หน้า viewer page (`/watch/:sessionId`)
- เรียก API เพื่อ get viewer token
- ใช้ LiveKit React SDK connect เข้า room
- subscribe + render remote video/audio track
- แสดงสถานะ (connecting, live, ended)
- auto-subscribe เมื่อมี track ใหม่

## Out of Scope
- transcript display (Phase 3)
- chat
- reactions

## Acceptance Criteria
- เข้าหน้า viewer ได้
- เห็น live video จาก publisher
- ได้ยิน audio จาก publisher
- แสดง connection status
- ถ้า publisher ยังไม่ live แสดงสถานะ waiting

## Files Likely Affected
- frontend/src/app/watch/[sessionId]/page.tsx
- frontend/src/components/Viewer.tsx

## Test Checklist
- [ ] เข้าหน้า /watch/:sessionId ได้
- [ ] เห็น video จาก publisher
- [ ] ได้ยิน audio
- [ ] publisher disconnect -> viewer เห็นสถานะเปลี่ยน

## Outcome


## Completion Evidence


## Completed At

