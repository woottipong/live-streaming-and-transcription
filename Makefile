.PHONY: dev-frontend dev-backend dev-worker

dev-frontend:
	cd frontend && pnpm dev

dev-backend:
	cd backend && env GOPROXY=https://proxy.golang.org,direct GOSUMDB=sum.golang.org go run github.com/air-verse/air@latest -c .air.toml

dev-worker:
	cd worker && env GOPROXY=https://proxy.golang.org,direct GOSUMDB=sum.golang.org go run github.com/air-verse/air@latest -c .air.toml
