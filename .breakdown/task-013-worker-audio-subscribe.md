# Task: Worker Audio Subscribe

## Task ID
task-013

## Epic
epic-07-transcription-worker

## Area
worker

## Status
todo

## Priority
high

## Depends On
- task-012

## Summary
Worker subscribe audio track จาก publisher ใน LiveKit room

## Scope
- auto-subscribe audio track เมื่อ publisher publish
- รับ audio frames จาก track
- convert audio เป็น format ที่ ASR ต้องการ (PCM 16-bit, 16kHz mono)
- ส่ง audio data ไปยัง channel สำหรับ ASR consumer
- handle track unpublished / publisher left

## Out of Scope
- video track
- multiple audio tracks
- noise reduction

## Acceptance Criteria
- worker subscribe audio track ได้เมื่อ publisher publish
- audio frames ถูกส่งเข้า channel
- audio format ถูกต้อง (PCM 16-bit, 16kHz mono)
- publisher leave → worker handle gracefully

## Files Likely Affected
- worker/internal/session/
- worker/internal/audio/

## Test Checklist
- [ ] publisher publish audio → worker รับ audio frames ได้
- [ ] log แสดง audio frame metadata (sample rate, channels)
- [ ] publisher leave → worker ไม่ crash

## Outcome


## Completion Evidence


## Completed At

