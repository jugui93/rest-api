package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jugui93/rest-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/fact", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)

	app.Get("/fact/:id", handlers.ShowFact)

	app.Patch("/fact/:id", handlers.UpdateFact)

	app.Delete("/fact/:id", handlers.DeleteFact)
}