# Live Streaming + Realtime Transcription Architecture

## Overview

ระบบนี้เป็นแพลตฟอร์มสำหรับ **Live Streaming พร้อม Realtime Transcription** โดยใช้ **LiveKit** เป็น media layer หลัก และมีองค์ประกอบดังนี้:

- **Admin Web** สำหรับสร้างและจัดการ stream session
- **Publisher** สำหรับส่ง audio/video เข้า stream
- **Viewer Web** สำหรับรับชม live stream และดู transcript แบบ realtime
- **Transcription Worker** สำหรับดึง audio จาก LiveKit แล้วส่งไปถอดเสียงกับ ASR provider
- **Backend API** สำหรับจัดการ session, metadata, token, webhook และ transcript delivery

เป้าหมายของเวอร์ชันแรกคือ:
- live stream ได้จริง
- ถอดเสียงได้จริงแบบ realtime
- viewer ดูวิดีโอและข้อความสดได้
- เก็บ final transcript ได้
- สถาปัตยกรรมไม่ซับซ้อนเกินจำเป็นสำหรับ MVP

---

# Goals

## In Scope
- สร้างและจัดการ live session ผ่าน admin web
- กำหนด metadata ของ session
- รองรับ publisher ผ่าน browser หรือ RTMP/OBS
- ใช้ LiveKit สำหรับ audio/video realtime
- ใช้ ASR provider 1 เจ้าใน MVP
- ถอดเสียงแบบ realtime
- ส่ง transcript ไปยัง viewer แบบ realtime ผ่าน WebSocket
- เก็บ final transcript ลง PostgreSQL

## Out of Scope
- Queue / background job system
- Multi-provider ASR fallback
- Subtitle export เช่น SRT/VTT
- AI summary / translation
- Advanced analytics
- Multi-worker orchestration ซับซ้อน
- Advanced speaker diarization

---

# High-Level Architecture

```text
[Admin Web / Viewer Web]
        |
        v
   [GoFiber API] ---------> [PostgreSQL]
      |     |
      |     +--> [LiveKit Server API]
      |     |
      |     +--> POST /start --> [Transcription Worker]
      |                                 |
      |     <-- POST /internal/transcripts/event --+
      |
      +--> [WebSocket Hub] --> [Viewer Web]

[Publisher Browser / OBS]
        |
        v
    [LiveKit Room]
        |
        v
[Transcription Worker] (joins as bot)
        |
        +--> subscribe audio track
        +--> stream audio to [ASR Provider]
        +--> POST transcript events to [GoFiber API]
```

---

# Technology Stack

## Frontend
- **Next.js**
- **LiveKit React SDK**
- Native **WebSocket** client

## Backend
- **Go**
- **Fiber** สำหรับ REST API และ WebSocket endpoint

## Media
- **LiveKit**
- รองรับ:
  - Browser publish
  - RTMP ingress ผ่าน OBS

## Data
- **PostgreSQL**
- **Redis** (optional ใน MVP, ใช้ภายหลังได้ถ้าต้องการ latest state หรือ multi-instance pub/sub)

## ASR
- ใช้ **ASR provider เดียว** ใน MVP
- ตัวอย่าง provider:
  - Deepgram
  - Google STT
  - AWS Transcribe
  - Azure Speech

---

# Core Components

## 1. Admin Web
หน้าที่:
- สร้าง session
- แก้ไข metadata
- ดูสถานะ session
- ได้ room info / ingest info สำหรับเริ่ม stream

ตัวอย่าง metadata:
- title
- description
- language
- ASR provider
- source type
- subtitle enabled

---

## 2. Viewer Web
หน้าที่:
- ดู live video/audio
- รับ transcript แบบ realtime
- แสดง subtitle overlay หรือ transcript panel

การเชื่อมต่อ:
- LiveKit room สำหรับ media
- WebSocket สำหรับ transcript

---

## 3. GoFiber API
หน้าที่:
- session CRUD
- auth/token generation
- LiveKit room/ingress orchestration
- webhook receiver จาก LiveKit
- transcript WebSocket endpoint
- รับ transcript event จาก worker
- persist final transcript

---

## 4. LiveKit
หน้าที่:
- รองรับ live audio/video แบบ realtime
- room management
- publisher/viewer connectivity
- ingress สำหรับ RTMP/OBS

LiveKit เป็น media plane หลักของระบบ

---

## 5. Transcription Worker
หน้าที่:
- join LiveKit room เป็น bot
- subscribe audio track
- stream audio ไปยัง ASR provider
- รับ partial/final transcript
- normalize transcript event
- ส่ง transcript event กลับไป backend
- persist final transcript หรือส่งให้ backend persist

---

## 6. PostgreSQL
หน้าที่:
- เก็บ session data
- เก็บ metadata
- เก็บ final transcript
- เก็บสถานะ session

---

# Realtime Flow

## 1. Create Session
Admin สร้าง live session จากหน้าเว็บ

### Flow
1. Admin กรอกข้อมูล session
2. Next.js เรียก API ไปที่ GoFiber
3. GoFiber:
   - validate input
   - save session ลง PostgreSQL
   - create LiveKit room
   - create ingress หากเป็น RTMP mode
4. API ส่ง session info กลับ frontend

### Output
- sessionId
- roomName
- ingressUrl / streamKey (ถ้ามี)
- publisher/viewer token endpoint

---

## 2. Start Live
รองรับ 2 รูปแบบ

### Browser Publisher
1. Publisher ขอ token จาก backend
2. Connect เข้า LiveKit room
3. Publish audio/video

### RTMP / OBS
1. OBS push stream ไปยัง LiveKit ingress
2. LiveKit รับ media เข้า room
3. Backend update สถานะ session เป็น `live`

---

## 3. Start Transcription
เมื่อ session เริ่ม live แล้ว

### Trigger
- LiveKit ส่ง webhook `TrackPublished` (audio track) มาที่ Backend
- Backend ตรวจว่า session มี `subtitle_enabled = true`
- ใช้ `TrackPublished` ไม่ใช่ `ParticipantJoined` เพราะ publisher อาจ join แล้วยังไม่ publish audio

### Flow
1. Backend รับ LiveKit webhook `TrackPublished`
2. Backend ตรวจว่า session เปิด subtitle
3. Backend POST `/start` ไปหา Transcription Worker พร้อม `{ sessionId, roomName, asrProvider, language }`
4. Backend update `transcription_status` เป็น `starting`
5. Worker join LiveKit room เป็น bot
6. Worker subscribe audio track
7. Worker connect ไป ASR provider แบบ streaming
8. Worker เริ่มส่ง audio chunk ไป ASR
9. Worker update `transcription_status` เป็น `running` ผ่าน backend

### Worker Control API
Worker เป็น long-running HTTP service รับ command จาก Backend:
- `POST /start` — เริ่ม transcription session
- `POST /stop`  — หยุด transcription, flush transcript ที่ค้าง

### Worker URL Config
Worker URL config ผ่าน environment variable: `WORKER_URL=http://worker:8081`

---

## 4. Realtime Transcript Delivery
### Flow
1. ASR ส่ง partial/final transcript กลับมา
2. Worker normalize ข้อมูลให้อยู่ใน schema กลาง
3. Worker POST `/internal/transcripts/event` ไปยัง Backend (per event)
4. Backend WS Hub broadcast event ไปทุก viewer ที่ subscribe sessionId นั้น
5. ถ้าเป็น `final` → Backend persist ลง PostgreSQL ด้วย

### Protocol: Worker → Backend
- **HTTP POST** `/internal/transcripts/event`
- ใช้ HTTP เพราะง่าย, stateless จาก worker มุมมอง
- Internal network latency < 1ms ไม่ใช่ปัญหาใน MVP
- อนาคตสามารถเปลี่ยนเป็น internal gRPC stream หรือ Redis pub/sub ได้

---

## 5. Viewer Playback
### Flow
1. Viewer เปิดหน้า watch
2. Frontend เรียก API เพื่อดึง session info
3. Viewer connect:
   - LiveKit room
   - WebSocket transcript endpoint
4. Frontend render:
   - live video
   - rolling transcript
   - subtitle overlay

---

## 6. End Session
### Flow
1. Publisher disconnect หรือ ingress stopped
2. Backend รับ LiveKit webhook (`ParticipantLeft` หรือ `IngressEnded`)
3. Update `stream_status` เป็น `ending`
4. Backend POST `/stop` ไปหา Worker
5. Worker flush transcript ที่ค้าง แล้ว disconnect จาก ASR และ LiveKit room
6. Backend update `stream_status` เป็น `ended`, `transcription_status` เป็น `stopped`
7. Backend ปิด LiveKit room (ถ้าจำเป็น)

---

# Realtime Data Flow

## Media Path
```text
Publisher / OBS
   -> LiveKit
   -> Viewer
```

## Transcript Path
```text
LiveKit audio
   -> Transcription Worker
   -> ASR Provider
   -> GoFiber API / WS Hub
   -> Viewer WebSocket
```

---

# Why No Queue in MVP

ใน MVP นี้ **ยังไม่ใช้ queue** ด้วยเหตุผลดังนี้:
- ลด complexity
- ทำ demo ให้เสร็จเร็ว
- ลดจำนวน moving parts
- โฟกัสที่ critical path ของ realtime system

Queue จะถูกเพิ่มภายหลังสำหรับงานประเภท:
- export subtitle
- summary
- translation
- analytics
- cleanup / archive

**หลักการสำคัญ:**  
realtime path ต้องสั้นที่สุดและไม่ผ่าน queue

---

# Realtime vs Async Boundary

## Realtime Path
ต้องเร็วและหน่วงต่ำ:
- stream ผ่าน LiveKit
- audio subscribe
- ASR streaming
- transcript push ผ่าน WebSocket

## Async Path
ยังไม่ทำใน MVP:
- subtitle export
- AI summary
- translation
- reporting
- archive

---

# API Design (MVP)

## Session APIs
- `POST /api/sessions`
- `GET /api/sessions/:id`
- `PATCH /api/sessions/:id`

## Token APIs
- `POST /api/sessions/:id/token/publisher`
- `POST /api/sessions/:id/token/viewer`

## Internal APIs
- `POST /internal/livekit/webhook`
- `POST /internal/transcripts/event`

## Worker Control APIs (Backend → Worker)
- `POST /start` — เริ่ม transcription: `{ sessionId, roomName, asrProvider, language }`
- `POST /stop`  — หยุด transcription และ flush: `{ sessionId }`

## WebSocket
- `GET /ws/transcripts/:sessionId`

---

# Data Model

## sessions
เก็บข้อมูล session หลัก

Fields:
- `id`
- `title`
- `description`
- `room_name`
- `ingress_id` — สำหรับ cleanup RTMP ingress เมื่อ session end
- `source_type`
- `language`
- `asr_provider`
- `subtitle_enabled`
- `stream_status`
- `transcription_status`
- `started_at`
- `ended_at`
- `created_at`
- `updated_at`

---

## transcript_segments
เก็บ final transcript เท่านั้น

Fields:
- `id`
- `session_id`
- `seq`
- `text`
- `start_time_ms`
- `end_time_ms`
- `speaker_id`
- `language`
- `confidence`
- `provider`
- `created_at`

---

# Transcript Event Schema

ตัวอย่าง event:

```json
{
  "sessionId": "sess_123",
  "seq": 101,
  "type": "partial",
  "text": "สวัสดีครับทุกท่าน",
  "startTimeMs": 1000,
  "endTimeMs": 2400,
  "speakerId": "speaker_1",
  "language": "th-TH",
  "confidence": 0.89,
  "isFinal": false
}
```

## Required Fields
- `sessionId`
- `seq`
- `type` (`partial` / `final`)
- `text`
- `startTimeMs`
- `endTimeMs`
- `language`
- `confidence`

## Notes
- partial ใช้แสดงผลสด และสามารถถูกแทนที่ได้
- final ใช้แสดงผลถาวร และถูกบันทึกลง DB
- `seq` ใช้เรียง event
- timestamp ใช้สำหรับ sync subtitle กับ stream

---

# Session State

แนะนำให้แยก state ออกเป็นหลายแกน

## stream_status
- `created`
- `ready`
- `waiting_for_input`
- `live`
- `ending`
- `ended`
- `failed`

## transcription_status
- `idle`
- `starting`
- `running`
- `stopping`
- `stopped`
- `failed`

เหตุผล:
- stream อาจ live ได้ แม้ transcription fail
- transcription อาจ reconnect แยกจาก media
- ทำ monitoring และ retry ได้ชัดกว่าใช้ state เดียว

---

# Deployment Strategy (MVP)

## Minimal Deployment
- Next.js deploy แยก
- GoFiber API deploy เป็น service เดียว
- Transcription Worker deploy แยกอีก service
- PostgreSQL เป็น managed database
- LiveKit ใช้ cloud service

## Optional
- Redis สำหรับ future scaling
- object storage สำหรับ recording/export ในอนาคต

---

# MVP Implementation Strategy

## Phase 1
- สร้าง admin page
- สร้าง viewer page
- ทำ session CRUD
- เชื่อม LiveKit room
- เชื่อม publisher / viewer flow

## Phase 2
- สร้าง transcription worker
- ดึง audio จาก room
- ส่งไป ASR
- รับ transcript partial/final

## Phase 3
- ทำ transcript websocket delivery
- แสดง transcript บน viewer page
- persist final transcript ลง DB

---

# Design Principles

## 1. Keep Realtime Path Short
อย่าให้เส้นทาง realtime ผ่าน component ที่ไม่จำเป็น

## 2. Separate Media and Transcript Concerns
media ใช้ LiveKit  
transcript ใช้ worker + WebSocket

## 3. Start with One ASR Provider
ลด complexity ใน MVP

## 4. Persist Only Final Transcript
partial transcript ไม่จำเป็นต้องลง DB ทุกครั้ง

## 5. Design for Future Extensibility
แม้ MVP ยังไม่มี queue แต่ควรออกแบบ service boundaries ให้รองรับอนาคตได้

---

# Future Extensions

สิ่งที่สามารถเพิ่มได้ภายหลัง:
- Queue / background jobs
- transcript export (SRT / VTT)
- translation
- summary
- analytics
- ASR fallback
- Redis pub/sub
- multi-instance websocket gateway
- advanced speaker diarization

---

# Final Summary

สถาปัตยกรรม MVP นี้ถูกออกแบบให้:
- เริ่มทำได้เร็ว
- realtime path ชัดเจน
- แยก concern ระหว่าง media, transcription, delivery
- รองรับการขยายในอนาคตได้
- ไม่ซับซ้อนเกินจำเป็นใน phase แรก

## MVP Stack Summary
- **Frontend:** Next.js
- **Backend:** Go + Fiber
- **Realtime Media:** LiveKit
- **Transcription Worker:** Go
- **Database:** PostgreSQL
- **Realtime Transcript Delivery:** WebSocket
- **Queue:** not included in MVP

## Core Principle
> Use LiveKit for media, use a Go worker for streaming ASR, use WebSocket for transcript delivery, and keep async/background concerns out of the MVP critical path.