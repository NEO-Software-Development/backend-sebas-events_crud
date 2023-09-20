package routes

import (
    "github.com/gofiber/fiber/v2"
    "crud-events-go/handlers"
)

func SetupEventRoutes(app *fiber.App) {
    eventRouter := app.Group("/events")

    eventRouter.Get("/", handlers.GetAllEvents)
    eventRouter.Get("/:id", handlers.GetEventByID)
    eventRouter.Post("/", handlers.CreateEvent)
    eventRouter.Put("/:id", handlers.UpdateEvent)
    eventRouter.Delete("/:id", handlers.DeleteEvent)
}