package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app  *fiber.App
	port int
}

func NewServer(port int) *Server {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100MB
		AppName:   "video-upload-service",
	})

	server := &Server{
		app:  app,
		port: port,
	}
	server.setRoute()

	return server
}

func (s *Server) Run() {
	s.app.Listen(fmt.Sprintf(":%d", s.port))
}

func (s *Server) setRoute() {
	s.app.Post("/upload", s.uploadVideo)
}
