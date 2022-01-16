package main

import (
	"log"
	"os"

	"github.com/PrinceNarteh/restaurant-management-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	}

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	router := app.Group("/api")
	routes.FoodRoutes(router)

	log.Fatal(app.Listen(port))
}
