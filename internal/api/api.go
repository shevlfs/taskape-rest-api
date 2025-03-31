package api

import (
	"fmt"
	"log"
	"taskape-rest-api/internal/api/handlers"
	"taskape-rest-api/internal/config"
	"taskape-rest-api/internal/grpc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app               *fiber.App
	config            *config.Config
	connectionManager *grpc.ConnectionManager
}

func NewServer(cfg *config.Config, connectionManager *grpc.ConnectionManager) *Server {
	app := fiber.New(fiber.Config{
		Prefork: true,
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
		DisableStartupMessage: !cfg.Debug,
	})

	// Add recovery middleware to prevent crashes
	app.Use(recover.New())

	return &Server{
		app:               app,
		config:            cfg,
		connectionManager: connectionManager,
	}
}

// connectionMiddleware ensures the gRPC connection is healthy before proceeding
func (s *Server) connectionMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client := s.connectionManager.GetClient()
		if client == nil || !client.CheckConnection() {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"success": false,
				"message": "Backend service temporarily unavailable, please try again shortly",
			})
		}
		return c.Next()
	}
}

func (s *Server) SetupRoutes() {
	// Apply connection middleware to all routes
	s.app.Use(s.connectionMiddleware())

	// Get the client for handlers
	client := s.connectionManager.GetClient()

	h := handlers.NewHandlers(client.BackendClient, s.config)
	SetupRoutes(s.app, h)
}

func (s *Server) Start() error {
	s.SetupRoutes()

	addr := fmt.Sprintf(":%s", s.config.Port)
	log.Printf("Starting server on %s", addr)
	return s.app.Listen(addr)
}
