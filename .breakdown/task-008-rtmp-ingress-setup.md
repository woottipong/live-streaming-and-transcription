# Task: RTMP Ingress Setup

## Task ID
task-008

## Epic
epic-04-publisher-flow

## Area
backend

## Status
todo

## Priority
medium

## Depends On
- task-004

## Summary
สร้าง RTMP ingress ผ่าน LiveKit สำหรับ OBS

## Scope
- สร้าง ingress เมื่อ session มี source_type = 'rtmp'
- save ingress_id ลง session record
- return ingress URL + stream key ใน session response
- ใช้ LiveKit Ingress API

## Out of Scope
- WHIP ingress
- HLS
- ingress deletion
- multiple ingress per session

## Acceptance Criteria
- สร้าง session ด้วย source_type = 'rtmp' แล้วได้ ingress URL + stream key
- ingress_id ถูก save ลง DB
- OBS สามารถ push stream เข้า ingress ได้

## Files Likely Affected
- backend/internal/livekit/ingress.go
- backend/internal/handler/session.go

## Test Checklist
- [ ] POST session ด้วย source_type=rtmp ได้ ingress URL
- [ ] ingress_id อยู่ใน DB
- [ ] OBS connect + push ได้จริง

## Outcome


## Completion Evidence


## Completed At

