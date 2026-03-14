package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"os"

	"github.com/iwoody/realtime-streaming/backend/internal/app"
	backendlivekit "github.com/iwoody/realtime-streaming/backend/internal/livekit"
	"github.com/iwoody/realtime-streaming/backend/internal/repository"
)

func main() {
	port := getEnv("BACKEND_PORT", "8080")
	databaseURL := getDatabaseURL()
	liveKitHTTPURL := getEnv("LIVEKIT_HTTP_URL", "http://livekit:7880")
	liveKitAPIKey := getEnv("LIVEKIT_API_KEY", "devkey")
	liveKitAPISecret := getEnv("LIVEKIT_API_SECRET", "secret")

	ctx := context.Background()
	pool, err := repository.NewPostgresPool(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	liveKitClient := backendlivekit.NewClient(liveKitHTTPURL, liveKitAPIKey, liveKitAPISecret)
	fiberApp := app.NewWithLiveKit(app.NewPostgresRepo(pool), liveKitClient)

	addr := ":" + port
	log.Printf("backend listening on %s", addr)

	if err := fiberApp.Listen(addr); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func getDatabaseURL() string {
	if value := os.Getenv("DATABASE_URL"); value != "" {
		return value
	}

	user := getEnv("POSTGRES_USER", "postgres")
	password := getEnv("POSTGRES_PASSWORD", "postgres")
	host := getEnv("POSTGRES_HOST", "postgres")
	port := getEnv("POSTGRES_PORT", "5432")
	database := getEnv("POSTGRES_DB", "realtime_streaming")

	databaseURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   net.JoinHostPort(host, port),
		Path:   database,
	}

	query := databaseURL.Query()
	query.Set("sslmode", "disable")
	databaseURL.RawQuery = query.Encode()

	return databaseURL.String()
}
