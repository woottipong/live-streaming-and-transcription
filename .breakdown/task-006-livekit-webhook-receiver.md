# Task: LiveKit Webhook Receiver

## Task ID
task-006

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
สร้าง endpoint รับ webhook จาก LiveKit เพื่อ update session state

## Scope
- POST /internal/livekit/webhook
- verify webhook signature ด้วย LiveKit SDK
- handle events: ParticipantJoined, TrackPublished, ParticipantLeft, IngressStarted, IngressEnded
- update stream_status ตาม event (เช่น TrackPublished -> live, ParticipantLeft -> ending)
- log events สำหรับ debug

## Out of Scope
- auto-start transcription (task-022)
- room cleanup

## Acceptance Criteria
- webhook endpoint รับ event ได้
- signature verification ทำงาน
- stream_status update ตาม event ที่ได้รับ
- event ที่ไม่รู้จักไม่ทำให้ crash

## Files Likely Affected
- backend/internal/handler/webhook.go
- backend/internal/livekit/webhook.go

## Test Checklist
- [ ] POST webhook ด้วย valid signature ผ่าน
- [ ] POST webhook ด้วย invalid signature ได้ 401
- [ ] TrackPublished -> stream_status เป็น live
- [ ] ParticipantLeft -> stream_status เป็น ending/ended
- [ ] unknown event ไม่ error

## Outcome


## Completion Evidence


## Completed At

