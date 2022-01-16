package routes

import (
	"github.com/PrinceNarteh/restaurant-management-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func InvoiceRoutes(router fiber.Router) {
	invoiceRouter := router.Group("/Invoices")
	invoiceRouter.Get("/", controllers.GetInvoices)
	invoiceRouter.Get("/:InvoiceId", controllers.GetInvoice)
	invoiceRouter.Post("/", controllers.CreateInvoice)
	invoiceRouter.Patch("/:InvoiceId", controllers.UpdateInvoice)
}
