# Project Status

## Current Phase
**Phase 1** — Backend foundation + Frontend skeleton + LiveKit integration

## Overall Progress
- Phase 1: 60% (task-001, task-002, task-003, task-004, task-005, and task-006 done)
- Phase 2: 0% (not started)
- Phase 3: 0% (not started)

## Current Priorities
1. ทำ publisher / viewer flow ผ่าน LiveKit
2. ทำ admin page สำหรับสร้าง/จัดการ session
3. ทำ worker scaffold
4. ทำ viewer media playback
5. ทำ RTMP ingress setup

## Blockers
- ยังไม่มี

---

## Epic Summary

- `✅ epic-01-project-setup` — โครงโปรเจกต์ backend, frontend, database, Docker พร้อมใช้งาน
- `✅ epic-02-session-management` — API และฐานข้อมูลสำหรับสร้าง, ดู, และแก้ไข session
- `🚧 epic-03-livekit-integration` — เชื่อม LiveKit สำหรับ room, token, webhook, และ ingress
- `📝 epic-04-publisher-flow` — รองรับ publisher ผ่าน browser และ RTMP/OBS
- `📝 epic-05-viewer-flow` — รองรับ viewer สำหรับดู live stream ผ่าน LiveKit
- `📝 epic-06-admin-web` — หน้า admin สำหรับสร้าง session, ดูสถานะ, และจัดการ room info
- `📝 epic-07-transcription-worker` — worker สำหรับ join room, subscribe audio, ส่งเข้า ASR, และ normalize transcript
- `📝 epic-08-transcript-delivery` — กระจาย transcript ผ่าน WebSocket และ persist final transcript
- `📝 epic-09-viewer-transcript-ui` — หน้า viewer สำหรับแสดง rolling transcript และ subtitle overlay

---

## Task Index

### Current Status Snapshot
- `✅ done`: task-001, task-002, task-003, task-004, task-005, task-006
- `🚧 in progress`: ยังไม่มี
- `🎯 next up`: task-007-publisher-browser-flow
- `⛔ blocked`: ยังไม่มี

### Done So Far
- `✅` Phase 1
  task-001-project-scaffold
  task-002-database-schema
  task-003-session-crud-api
  task-004-livekit-room-creation
  task-005-token-generation-api
  task-006-livekit-webhook-receiver

### Ready To Start
- `🎯` Phase 1
  task-007-publisher-browser-flow
  task-008-rtmp-ingress-setup
  task-009-viewer-media-playback
  task-010-admin-session-page

### Backlog
- `📝` Phase 2
  task-011-worker-scaffold
  task-012-worker-livekit-join
  task-013-worker-audio-subscribe
  task-014-worker-asr-streaming
  task-015-worker-control-api
  task-016-worker-transcript-normalize
- `📝` Phase 3
  task-017-transcript-ws-hub
  task-018-transcript-event-receiver
  task-019-transcript-persist
  task-020-viewer-transcript-ws-client
  task-021-viewer-subtitle-overlay
  task-022-auto-start-transcription

### Phase 1 — Backend + Frontend + LiveKit

| Task | Area | Priority | Status |
|------|------|----------|--------|
| task-001-project-scaffold | infra | high | ✅ done |
| task-002-database-schema | backend | high | ✅ done |
| task-003-session-crud-api | backend | high | ✅ done |
| task-004-livekit-room-creation | backend | high | ✅ done |
| task-005-token-generation-api | backend | high | ✅ done |
| task-006-livekit-webhook-receiver | backend | high | ✅ done |
| task-007-publisher-browser-flow | frontend | high | 📝 todo |
| task-008-rtmp-ingress-setup | backend | medium | 📝 todo |
| task-009-viewer-media-playback | frontend | high | 📝 todo |
| task-010-admin-session-page | frontend | high | 📝 todo |

### Phase 2 — Transcription Worker

| Task | Area | Priority | Status |
|------|------|----------|--------|
| task-011-worker-scaffold | worker | high | 📝 todo |
| task-012-worker-livekit-join | worker | high | 📝 todo |
| task-013-worker-audio-subscribe | worker | high | 📝 todo |
| task-014-worker-asr-streaming | worker | high | 📝 todo |
| task-015-worker-control-api | worker | high | 📝 todo |
| task-016-worker-transcript-normalize | worker | medium | 📝 todo |

### Phase 3 — Transcript Delivery + UI

| Task | Area | Priority | Status |
|------|------|----------|--------|
| task-017-transcript-ws-hub | backend | high | 📝 todo |
| task-018-transcript-event-receiver | backend | high | 📝 todo |
| task-019-transcript-persist | backend | high | 📝 todo |
| task-020-viewer-transcript-ws-client | frontend | high | 📝 todo |
| task-021-viewer-subtitle-overlay | frontend | medium | 📝 todo |
| task-022-auto-start-transcription | backend | high | 📝 todo |

---

## Definition of Done

task จะถือว่า done เมื่อ:
1. code ทำงานได้ตาม acceptance criteria ทุกข้อ
2. ไม่มี compilation error
3. ทดสอบ manual flow ผ่าน (ตาม test checklist ใน task file)
4. ไม่ทำลาย task อื่นที่ done แล้ว
5. code อยู่ใน scope ที่กำหนด ไม่ขยายงานเกิน

## Working Rules

1. **ทำทีละ task** — อย่าข้าม task หรือทำหลาย task พร้อมกัน
2. **อ้างอิง `docs/ARCHITECTURE.md` เสมอ** — เป็น source of truth ด้านสถาปัตยกรรม
3. **อย่าขยาย scope** — ถ้า task บอกว่า out of scope ห้ามทำ
4. **update status ทุกครั้ง** — เมื่อ task เปลี่ยนสถานะ ให้ update ทั้ง task file และไฟล์นี้
5. **ห้ามทำนอก MVP scope** — ไม่ทำ queue, multi-ASR, export, summary, translation
6. **ถ้าต้องสมมติอะไร ให้ระบุชัด** — เช่น ASR provider ที่เลือก, LiveKit config
