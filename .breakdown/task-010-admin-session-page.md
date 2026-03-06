# Task: Admin Session Page

## Task ID
task-010

## Epic
epic-06-admin-web

## Area
frontend

## Status
todo

## Priority
high

## Depends On
- task-003

## Summary
ทำหน้า admin สำหรับสร้างและดู session

## Scope
- หน้า admin create session (`/admin/sessions/new`)
- form สำหรับกรอก: title, description, language, source_type, asr_provider, subtitle_enabled
- เรียก POST /api/sessions
- หน้า admin session detail (`/admin/sessions/:id`)
- แสดง session info, stream_status, transcription_status, room_name
- แสดง ingress URL/stream key (ถ้ามี)
- link ไปหน้า publisher และ viewer

## Out of Scope
- session list page
- edit form
- delete
- authentication

## Acceptance Criteria
- สร้าง session จาก form ได้
- หลังสร้างเสร็จ redirect ไปหน้า detail
- หน้า detail แสดง session info ครบ
- มี link ไป /publish/:id และ /watch/:id

## Files Likely Affected
- frontend/src/app/admin/sessions/new/page.tsx
- frontend/src/app/admin/sessions/[id]/page.tsx

## Test Checklist
- [ ] สร้าง session จาก form ได้
- [ ] redirect ไปหน้า detail
- [ ] หน้า detail แสดง info ครบ
- [ ] link publisher/viewer ใช้ได้

## Outcome


## Completion Evidence


## Completed At

