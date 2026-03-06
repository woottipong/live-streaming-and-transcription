# Task: Viewer Subtitle Overlay

## Task ID
task-021

## Epic
epic-09-viewer-transcript-ui

## Area
frontend

## Status
todo

## Priority
medium

## Depends On
- task-020

## Summary
แสดง subtitle overlay บน video player จาก latest transcript

## Scope
- overlay component แสดงบน video area
- แสดง latest partial หรือ final text
- fade out เมื่อไม่มี text ใหม่ (เช่น 3 วินาที)
- toggle on/off subtitle
- styling: semi-transparent background, readable font

## Out of Scope
- multi-line subtitle
- positioning options
- font size settings
- SRT/VTT

## Acceptance Criteria
- เห็น subtitle overlay บน video
- แสดง latest transcript text
- text fade out หลัง 3 วินาทีไม่มี update
- toggle on/off ทำงาน

## Files Likely Affected
- frontend/src/components/SubtitleOverlay.tsx

## Test Checklist
- [ ] เห็น subtitle บน video
- [ ] text update เมื่อมี transcript ใหม่
- [ ] fade out หลัง 3 วินาที
- [ ] toggle off → ไม่เห็น subtitle

## Outcome


## Completion Evidence


## Completed At

