package router

import (
	"github.com/gofiber/fiber"
	"github.com/tdwiid/amphibi/handler"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// routes
	app.Get("/", handler.GetAllProducts)
	app.Get("/:id", handler.GetSingleProduct)
	app.Post("/", handler.CreateProduct)
	app.Delete("/:id", handler.DeleteProduct)

}
