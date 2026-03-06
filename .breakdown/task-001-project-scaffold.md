# Task: Project Scaffold

## Task ID
task-001

## Epic
epic-01-project-setup

## Area
infra

## Status
todo

## Priority
high

## Depends On
- none

## Summary
ตั้งโครงสร้างโปรเจกต์ monorepo สำหรับ backend (Go + Fiber), frontend (Next.js), worker (Go), docker-compose สำหรับ dev

## Scope
- สร้าง Go module สำหรับ backend (`backend/`)
- สร้าง Go module สำหรับ worker (`worker/`)
- สร้าง Next.js project สำหรับ frontend (`frontend/`)
- สร้าง `docker-compose.yml` สำหรับ PostgreSQL + LiveKit (dev)
- สร้าง `.env.example` สำหรับ config
- backend ต้อง run ได้ (hello world endpoint)
- frontend ต้อง run ได้ (default Next.js page)
- worker ต้อง run ได้ (hello world log)

## Out of Scope
- CI/CD
- production deployment
- Dockerfile สำหรับ app

## Acceptance Criteria
- `go run ./backend` เริ่มได้ ฟัง port 8080
- `go run ./worker` เริ่มได้ ฟัง port 8081
- `npm run dev` ใน frontend เริ่มได้
- `docker compose up` รัน PostgreSQL + LiveKit ได้
- `.env.example` มี config ครบ (DB, LiveKit, Worker URL)

## Files Likely Affected
- backend/main.go
- backend/go.mod
- worker/main.go
- worker/go.mod
- frontend/package.json
- docker-compose.yml
- .env.example

## Test Checklist
- [ ] `go run ./backend` เริ่มได้ไม่ error
- [ ] `go run ./worker` เริ่มได้ไม่ error
- [ ] `npm run dev` ใน frontend เริ่มได้
- [ ] `docker compose up -d` รัน postgres ได้
- [ ] `.env.example` มีค่าครบ

## Outcome


## Completion Evidence


## Completed At

