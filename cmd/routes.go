package main

import (
	"github.com/celiovjunior/go-api-docker/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
}
