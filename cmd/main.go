package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jugui93/rest-api/database"
	"github.com/jugui93/rest-api/handlers"
)

func main() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	database.DB.Connect(dsn)
    app := fiber.New()
	handlers := handlers.NewHandlers()

	SetupRoutes(app, handlers)

    app.Listen(":3000")
}