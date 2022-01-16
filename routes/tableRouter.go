package routes

import (
	"github.com/PrinceNarteh/restaurant-management-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func TableRoutes(router fiber.Router) {
	tableRouter := router.Group("/tables")
	tableRouter.Get("/", controllers.GetTables)
	tableRouter.Get("/:tableId", controllers.GetTable)
	tableRouter.Post("/", controllers.CreateTable)
	tableRouter.Patch("/:tableId", controllers.UpdateTable)
}
