package controllers

import (
	"github.com/KadirOzerOzturk/procguard-agent/app/services"
	"github.com/gofiber/fiber/v2"
)

func CollectSystemStats(c *fiber.Ctx) error {
	stats, err := services.CollectSystemStats()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(stats)
}

func SendStats(c *fiber.Ctx) error {
	stats, err := services.CollectSystemStats()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.SendStatsToAPI(stats); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "✅ Stats sent to backend"})
}
