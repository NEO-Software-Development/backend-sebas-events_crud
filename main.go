package main

import (
    "github.com/gofiber/fiber/v2"
    "crud-events-go/routes"
)

func main() {
    app := fiber.New()

    // Configura las rutas de eventos
    routes.SetupEventRoutes(app)

    // Ejecuta la aplicación en el puerto 3000
    app.Listen(":3000")
}