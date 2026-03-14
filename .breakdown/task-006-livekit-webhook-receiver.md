# Task: LiveKit Webhook Receiver

## Task ID
task-006

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
- [x] POST webhook ด้วย valid signature ผ่าน
- [x] POST webhook ด้วย invalid signature ได้ 401
- [x] TrackPublished -> stream_status เป็น live
- [x] ParticipantLeft -> stream_status เป็น ending/ended
- [x] unknown event ไม่ error

## Outcome
เพิ่ม `POST /internal/livekit/webhook` ที่ตรวจ signature ด้วย LiveKit auth SDK, parse payload เป็น webhook event, และ map event ไปอัปเดต `stream_status` ตาม `room_name`

รองรับ event สำคัญใน MVP ได้แก่ `ParticipantJoined`, `TrackPublished`, `ParticipantLeft`, `IngressStarted`, และ `IngressEnded` พร้อม log event ที่รับเข้าและ ignore event ที่ไม่รู้จักอย่างปลอดภัย

## Completion Evidence
- รัน `env GOCACHE=/tmp/go-cache go test ./...` ใน `backend/` ผ่าน
- Automated tests ครอบคลุม valid signature, invalid signature (`401`), `track_published -> live`, `participant_left -> ending`, และ unknown event -> ignored
- webhook route ตอบกลับได้โดยไม่ crash เมื่อเจอ event ที่ไม่อยู่ใน scope

## Completed At
2026-03-15 02:35:00 +07
