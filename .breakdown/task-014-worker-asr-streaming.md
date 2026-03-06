# Task: Worker ASR Streaming

## Task ID
task-014

## Epic
epic-07-transcription-worker

## Area
worker

## Status
todo

## Priority
high

## Depends On
- task-013

## Summary
Worker ส่ง audio stream ไป ASR provider แล้วรับ partial/final transcript กลับมา

## Scope
- implement ASR client interface
- implement 1 ASR provider (เลือก Deepgram เป็น MVP provider)
- stream audio chunks ไป ASR แบบ realtime
- รับ partial/final transcript จาก ASR
- config: ASR_API_KEY, ASR_PROVIDER จาก env
- handle ASR connection error / reconnect

## Out of Scope
- multi-provider
- fallback
- custom model
- speaker diarization

## Acceptance Criteria
- audio ถูกส่งไป ASR provider ได้
- ได้รับ partial transcript กลับมา
- ได้รับ final transcript กลับมา
- ASR disconnect → reconnect attempt

## Files Likely Affected
- worker/internal/asr/
- worker/internal/asr/deepgram/

## Test Checklist
- [ ] stream audio → ได้ partial transcript กลับมา
- [ ] ได้ final transcript กลับมา
- [ ] ASR config ผิด → error message ชัดเจน

## Outcome


## Completion Evidence


## Completed At

