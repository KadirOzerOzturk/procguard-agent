package controllers

import (
	"github.com/KadirOzerOzturk/procguard-agent/app/services"
	"github.com/gofiber/fiber/v2"
)

type KillRequest struct {
	Pid int32 `json:"pid"`
}

func KillProcess(c *fiber.Ctx) error {
	var req KillRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := services.KillProcess(req.Pid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Process killed successfully"})
}

func GetTopProcesses(c *fiber.Ctx) error {
	topProcs, err := services.GetTopProcesses(5)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(topProcs)
}
func GetAllProcesses(c *fiber.Ctx) error {
	allProcs, err := services.GetAllProcesses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(allProcs)
}
