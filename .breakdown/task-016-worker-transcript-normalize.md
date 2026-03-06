# Task: Worker Transcript Normalize

## Task ID
task-016

## Epic
epic-07-transcription-worker

## Area
worker

## Status
todo

## Priority
medium

## Depends On
- task-014

## Summary
Normalize transcript event จาก ASR provider ให้อยู่ใน schema กลางตาม ARCHITECTURE.md แล้วส่งไป Backend

## Scope
- normalize ASR response เป็น transcript event schema: sessionId, seq, type, text, startTimeMs, endTimeMs, speakerId, language, confidence, isFinal
- generate seq number ต่อเนื่องต่อ session
- POST ไปที่ Backend `/internal/transcripts/event`
- handle both partial and final events
- config: BACKEND_URL จาก env

## Out of Scope
- batch send
- retry queue
- local persistence

## Acceptance Criteria
- transcript event ถูก normalize ตาม schema
- seq เพิ่มขึ้นต่อเนื่อง
- partial event ถูกส่งไป backend
- final event ถูกส่งไป backend
- POST failure ถูก log (ไม่ block pipeline)

## Files Likely Affected
- worker/internal/transcript/
- worker/internal/sender/

## Test Checklist
- [ ] ASR partial → normalized event ถูก POST ไป backend
- [ ] ASR final → normalized event ถูก POST ไป backend
- [ ] seq เพิ่มขึ้นต่อเนื่อง
- [ ] backend down → log error ไม่ crash

## Outcome


## Completion Evidence


## Completed At

