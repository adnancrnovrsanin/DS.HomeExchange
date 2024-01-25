package routes

import (
	"DS.HomeExchange.Homes/controllers"
	"github.com/gofiber/fiber/v2"
)

func HomeRoutes(app *fiber.App) {
	app.Get("/", controllers.GetAllHomes)
	app.Post("/", controllers.CreateHome)
	app.Get("/:id", controllers.GetHome)
	app.Put("/", controllers.UpdateHome)
	app.Get("/owner/:id", controllers.GetHomesByOwner)
	app.Delete("/:id", controllers.DeleteHome)
	app.Post("/:id/reject", controllers.RejectHome)
	app.Post("/:id/accept", controllers.PerformHomeExchange)
}
