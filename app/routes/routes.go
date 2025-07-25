package routes

import (
	"github.com/KadirOzerOzturk/procguard-agent/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	agentGroup := app.Group("/agents")
	agentGroup.Get("/stats", controllers.CollectSystemStats)
	agentGroup.Post("/send", controllers.SendStats)
	agentGroup.Get("/ping", controllers.Ping)
	agentGroup.Get("/info", controllers.GetAgentInfo)

	processGroup := app.Group("/processes")
	processGroup.Post("/kill/:pid", controllers.KillProcess)
	processGroup.Get("/top", controllers.GetTopProcesses)
	processGroup.Get("/all", controllers.GetAllProcesses)
}
