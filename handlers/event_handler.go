package handlers

import (
    "github.com/gofiber/fiber/v2"
    "crud-events-go/models"
)

var events []models.Event

// GetAllEvents devuelve todos los eventos
func GetAllEvents(c *fiber.Ctx) error {
    return c.JSON(events)
}

// GetEventByID devuelve un evento por su ID
func GetEventByID(c *fiber.Ctx) error {
    // Lógica para buscar un evento por su ID y devolverlo como JSON

    // Aquí debes retornar una respuesta apropiada
    return nil
}

// CreateEvent crea un nuevo evento
func CreateEvent(c *fiber.Ctx) error {
    // Lógica para crear un nuevo evento y agregarlo a la lista de eventos

    // Aquí debes retornar una respuesta apropiada
    return nil
}

// UpdateEvent actualiza un evento existente
func UpdateEvent(c *fiber.Ctx) error {
    // Lógica para actualizar un evento existente por su ID

    // Aquí debes retornar una respuesta apropiada
    return nil
}

// DeleteEvent elimina un evento por su ID
func DeleteEvent(c *fiber.Ctx) error {
    // Lógica para eliminar un evento por su ID

    // Aquí debes retornar una respuesta apropiada
    return nil
}