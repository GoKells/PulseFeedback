package server

import (
	"github.com/gofiber/fiber/v2"

	"PulseFeedback/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "PulseFeedback",
			AppName:      "PulseFeedback",
		}),

		db: database.New(),
	}

	return server
}
