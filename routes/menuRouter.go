package routes

import (
	"github.com/PrinceNarteh/restaurant-management-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func MenuRoutes(router fiber.Router) {
	menuRouter := router.Group("/menus")
	menuRouter.Get("/", controllers.GetMenus)
	menuRouter.Get("/:menuId", controllers.GetMenu)
	menuRouter.Post("/", controllers.CreateMenu)
	menuRouter.Patch("/:menuId", controllers.UpdateMenu)
}
