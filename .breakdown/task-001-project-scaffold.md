# Task: Project Scaffold

## Task ID
task-001

## Epic
epic-01-project-setup

## Area
infra

## Status
done

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
- `pnpm dev` ใน frontend เริ่มได้
- `docker compose up` รัน PostgreSQL + LiveKit ได้
- `.env.example` มี config ครบ (DB, LiveKit, Worker URL)

## Files Likely Affected
- backend/main.go
- backend/go.mod
- worker/main.go
- worker/go.mod
- frontend/package.json
- frontend/components.json
- frontend/components/ui/button.jsx
- frontend/components/ui/card.jsx
- frontend/components/ui/badge.jsx
- docker-compose.yml
- .env.example

## Test Checklist
- [x] `go run ./backend` เริ่มได้ไม่ error
- [x] `go run ./worker` เริ่มได้ไม่ error
- [x] `pnpm dev` ใน frontend เริ่มได้
- [x] `docker compose up -d` รัน postgres ได้
- [x] `.env.example` มีค่าครบ

## Outcome
สร้าง monorepo scaffold สำหรับ backend, worker, frontend, `docker-compose.yml`, และ `.env.example` แล้ว พร้อม verify การรันของ backend, worker, และ frontend จาก root repository

frontend ถูกตั้งค่าให้ใช้ `pnpm` และ initialize `shadcn/ui` เรียบร้อย พร้อมหน้าแรกที่ compose ด้วย `Button`, `Card`, และ `Badge`

## Completion Evidence
`env GOPROXY=https://proxy.golang.org,direct GOSUMDB=sum.golang.org go run ./backend` ฟังที่ `:8080` และตอบ `/health`
`env GOPROXY=https://proxy.golang.org,direct GOSUMDB=sum.golang.org go run ./worker` ฟังที่ `:8081` และตอบ `/health`
`pnpm install` สร้าง `frontend/pnpm-lock.yaml` และ `pnpm dev` ใน `frontend/` ขึ้นที่ `http://localhost:3000`
`pnpm build` ใน `frontend/` ผ่าน และหน้า `/` render shadcn UI ใหม่ได้
`docker compose up -d` ขึ้น PostgreSQL และ LiveKit สำเร็จ โดย PostgreSQL publish ที่ `5433` เพื่อหลบพอร์ต `5432` ที่ถูกใช้อยู่บนเครื่อง และ `pg_isready -U postgres -d realtime_streaming` ตอบรับ connection


## Completed At
2026-03-14 23:48:30 +07
