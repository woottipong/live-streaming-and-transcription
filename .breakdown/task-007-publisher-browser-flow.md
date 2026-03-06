# Task: Publisher Browser Flow

## Task ID
task-007

## Epic
epic-04-publisher-flow

## Area
frontend

## Status
todo

## Priority
high

## Depends On
- task-005

## Summary
ทำหน้า publisher ใน frontend ให้ connect LiveKit room แล้ว publish audio/video

## Scope
- หน้า publisher page (`/publish/:sessionId`)
- เรียก API เพื่อ get publisher token
- ใช้ LiveKit React SDK connect เข้า room
- publish camera + microphone
- แสดง local video preview
- แสดงสถานะ connection (connecting, connected, disconnected)
- ปุ่ม start/stop publish

## Out of Scope
- screen share
- multiple cameras
- chat
- recording

## Acceptance Criteria
- เข้าหน้า publisher ได้
- กด start แล้ว publish audio/video เข้า LiveKit room
- เห็น local video preview
- กด stop แล้วหยุด publish
- แสดง connection status

## Files Likely Affected
- frontend/src/app/publish/[sessionId]/page.tsx
- frontend/src/components/Publisher.tsx

## Test Checklist
- [ ] เข้าหน้า /publish/:sessionId ได้
- [ ] กด start -> เห็น video preview
- [ ] ใน LiveKit dashboard เห็น participant ใหม่
- [ ] กด stop -> หยุด publish

## Outcome


## Completion Evidence


## Completed At

