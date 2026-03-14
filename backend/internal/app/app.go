package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iwoody/realtime-streaming/backend/internal/handler"
	"github.com/iwoody/realtime-streaming/backend/internal/model"
	"github.com/iwoody/realtime-streaming/backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Session = model.Session
type CreateSessionParams = repository.CreateSessionParams
type UpdateSessionParams = repository.UpdateSessionParams

var ErrSessionNotFound = repository.ErrSessionNotFound

type SessionRepository interface {
	Create(context.Context, CreateSessionParams) (Session, error)
	GetByID(context.Context, uuid.UUID) (Session, error)
	Update(context.Context, uuid.UUID, UpdateSessionParams) (Session, error)
}

func New(repo SessionRepository) *fiber.App {
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

func NewPostgresRepo(pool *pgxpool.Pool) SessionRepository {
	return repository.NewPostgresSessionRepository(pool)
}
