# Task: Viewer Transcript WebSocket Client

## Task ID
task-020

## Epic
epic-09-viewer-transcript-ui

## Area
frontend

## Status
todo

## Priority
high

## Depends On
- task-009
- task-017

## Summary
Viewer page connect WebSocket เพื่อรับ transcript events แบบ realtime

## Scope
- connect WebSocket ที่ /ws/transcripts/:sessionId
- parse transcript event JSON
- เก็บ transcript state (list of segments)
- partial event → update/replace segment ที่มี seq เดียวกัน
- final event → add permanent segment
- แสดง rolling transcript list
- auto-scroll ไปล่าสุด
- handle reconnect เมื่อ WS disconnect

## Out of Scope
- subtitle overlay (task-021)
- styling
- search
- export

## Acceptance Criteria
- viewer เห็น transcript events แบบ realtime
- partial events ถูก replace เมื่อมี final
- transcript list เรียงตาม seq
- WS disconnect → reconnect อัตโนมัติ

## Files Likely Affected
- frontend/src/components/TranscriptPanel.tsx
- frontend/src/hooks/useTranscriptWS.ts

## Test Checklist
- [ ] เห็น transcript ปรากฏแบบ realtime
- [ ] partial → final replacement ทำงาน
- [ ] auto-scroll ทำงาน
- [ ] disconnect → reconnect ได้

## Outcome


## Completion Evidence


## Completed At

