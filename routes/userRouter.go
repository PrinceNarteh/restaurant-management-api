package routes

import (
	controller "github.com/PrinceNarteh/restaurant-management-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(incomingRoute fiber.Router) {
	router := incomingRoute.Group("/users")
	router.Get("/", controller.GetUsers)
	router.Get("/:userId", controller.GetUser)
	router.Post("/register", controller.Register)
	router.Post("/login", controller.Login)
}
