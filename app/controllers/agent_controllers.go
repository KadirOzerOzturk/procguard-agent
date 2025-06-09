package controllers

import (
	"github.com/KadirOzerOzturk/procguard-agent/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetAgentInfo(c *fiber.Ctx) error {
	agentInfo, err := services.GetAgentInfo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(agentInfo)
}
