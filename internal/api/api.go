package api

import (
	"fmt"
	"log"
	"taskape-rest-api/internal/api/handlers"
	"taskape-rest-api/internal/config"
	proto "taskape-rest-api/proto"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app    *fiber.App
	config *config.Config
	client proto.BackendRequestsClient
}

func NewServer(cfg *config.Config, client proto.BackendRequestsClient) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			if cfg.Debug {
				log.Printf("Error: %v", err)
			}

			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	return &Server{
		app:    app,
		config: cfg,
		client: client,
	}
}

func (s *Server) SetupRoutes() {
	h := handlers.NewHandlers(s.client, s.config)
	SetupRoutes(s.app, h)
}

func (s *Server) Start() error {
	s.SetupRoutes()

	addr := fmt.Sprintf(":%s", s.config.Port)
	log.Printf("Starting server on %s", addr)
	return s.app.Listen(addr)
}
