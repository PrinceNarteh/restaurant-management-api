package routes

import (
	"github.com/PrinceNarteh/restaurant-management-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func InvoiceRoutes(router fiber.Router) {
	invoiceRouter := router.Group("/invoices")
	invoiceRouter.Get("/", controllers.GetInvoices)
	invoiceRouter.Get("/:invoiceId", controllers.GetInvoice)
	invoiceRouter.Post("/", controllers.CreateInvoice)
	invoiceRouter.Patch("/:invoiceId", controllers.UpdateInvoice)
}
