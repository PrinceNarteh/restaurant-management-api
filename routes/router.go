package routes

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	router := app.Group("/api")
	FoodRoutes(router)
	InvoiceRoutes(router)
	MenuRoutes(router)
	TableRoutes(router)
	UserRoutes(router)
}
