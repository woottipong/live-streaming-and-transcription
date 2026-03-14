# Project Status

## Current Phase
**Phase 1** — Backend foundation + Frontend skeleton + LiveKit integration

## Overall Progress
- Phase 1: 10% (task-001 done)
- Phase 2: 0% (not started)
- Phase 3: 0% (not started)

## Current Priorities
1. ทำ session CRUD API + database schema
2. เชื่อม LiveKit room creation
3. ทำ publisher / viewer flow ผ่าน LiveKit
4. ทำ admin page สำหรับสร้าง/จัดการ session
5. เตรียม worker integration บน frontend shell ที่ใช้ shadcn/ui แล้ว

## Blockers
- ยังไม่มี

---

## Epic Summary

### epic-01-project-setup
ตั้งโครงสร้างโปรเจกต์ backend, frontend, database, Docker

### epic-02-session-management
Session CRUD, database schema, session state management

### epic-03-livekit-integration
LiveKit room creation, token generation, webhook receiver, ingress setup

### epic-04-publisher-flow
Publisher connect ผ่าน browser + RTMP/OBS ingress

### epic-05-viewer-flow
Viewer connect เข้า LiveKit room ดู live stream

### epic-06-admin-web
Admin page สำหรับสร้าง session, ดูสถานะ, แสดง room info

### epic-07-transcription-worker
Worker service: join room, subscribe audio, stream ไป ASR, normalize transcript

### epic-08-transcript-delivery
WebSocket hub, transcript broadcast ไป viewer, persist final transcript

### epic-09-viewer-transcript-ui
Viewer page แสดง rolling transcript + subtitle overlay

---

## Task Index

### Phase 1 — Backend + Frontend + LiveKit

| Task | Epic | Area | Status | Priority |
|------|------|------|--------|----------|
| task-001-project-scaffold | epic-01 | infra | done | high |
| task-002-database-schema | epic-02 | backend | todo | high |
| task-003-session-crud-api | epic-02 | backend | todo | high |
| task-004-livekit-room-creation | epic-03 | backend | todo | high |
| task-005-token-generation-api | epic-03 | backend | todo | high |
| task-006-livekit-webhook-receiver | epic-03 | backend | todo | high |
| task-007-publisher-browser-flow | epic-04 | frontend | todo | high |
| task-008-rtmp-ingress-setup | epic-04 | backend | todo | medium |
| task-009-viewer-media-playback | epic-05 | frontend | todo | high |
| task-010-admin-session-page | epic-06 | frontend | todo | high |

### Phase 2 — Transcription Worker

| Task | Epic | Area | Status | Priority |
|------|------|------|--------|----------|
| task-011-worker-scaffold | epic-07 | worker | todo | high |
| task-012-worker-livekit-join | epic-07 | worker | todo | high |
| task-013-worker-audio-subscribe | epic-07 | worker | todo | high |
| task-014-worker-asr-streaming | epic-07 | worker | todo | high |
| task-015-worker-control-api | epic-07 | worker | todo | high |
| task-016-worker-transcript-normalize | epic-07 | worker | todo | medium |

### Phase 3 — Transcript Delivery + UI

| Task | Epic | Area | Status | Priority |
|------|------|------|--------|----------|
| task-017-transcript-ws-hub | epic-08 | backend | todo | high |
| task-018-transcript-event-receiver | epic-08 | backend | todo | high |
| task-019-transcript-persist | epic-08 | backend | todo | high |
| task-020-viewer-transcript-ws-client | epic-09 | frontend | todo | high |
| task-021-viewer-subtitle-overlay | epic-09 | frontend | todo | medium |
| task-022-auto-start-transcription | epic-03 | backend | todo | high |

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
