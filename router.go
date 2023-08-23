package main

import (
	"iwexlmsapi/routes/auth"
	"iwexlmsapi/routes/branch"
	"iwexlmsapi/routes/class"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func setupRoutes(app *fiber.App) {
	mainRouter := app.Group("/api/v2", logger.New())
	auth.SetupAuthRoute(&mainRouter)
	branch.SetupBranchRoutes(&mainRouter)
	class.SetupClassRoutes(&mainRouter)
}
