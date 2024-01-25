package main

import (
	"DS.HomeExchange.Homes/configs"
	"DS.HomeExchange.Homes/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	configs.ConnectDB()

	routes.HomeRoutes(app)

	app.Listen(":8080")
}
