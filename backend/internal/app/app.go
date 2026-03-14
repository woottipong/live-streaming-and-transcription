package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwoody/realtime-streaming/backend/internal/handler"
	"github.com/iwoody/realtime-streaming/backend/internal/model"
	"github.com/iwoody/realtime-streaming/backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Type aliases for test convenience.
type Session = model.Session
type CreateSessionParams = repository.CreateSessionParams
type UpdateSessionParams = repository.UpdateSessionParams

var ErrSessionNotFound = repository.ErrSessionNotFound

func New(repo repository.SessionRepository) *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "backend",
			"message": "hello world",
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	handler.NewSessionHandler(repo).Register(app)

	return app
}

func NewPostgresRepo(pool *pgxpool.Pool) repository.SessionRepository {
	return repository.NewPostgresSessionRepository(pool)
}
