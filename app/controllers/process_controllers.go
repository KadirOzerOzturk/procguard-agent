package controllers

import (
	"strconv"

	"github.com/KadirOzerOzturk/procguard-agent/app/services"
	"github.com/gofiber/fiber/v2"
)

type KillRequest struct {
	Pid int32 `json:"pid"`
}

func KillProcess(c *fiber.Ctx) error {

	var req KillRequest
	pidStr := c.Params("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid PID"})
	}
	req.Pid = int32(pid)
	if req.Pid < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid PID"})
	}

	if req.Pid == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "PID is required"})
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
