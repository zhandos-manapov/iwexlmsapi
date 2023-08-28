package main

import (
	"iwexlmsapi/routes/auth"
	"iwexlmsapi/routes/city"
	"iwexlmsapi/routes/course"
	"iwexlmsapi/routes/files"
	"iwexlmsapi/routes/level"

	"iwexlmsapi/routes/auth"
	"iwexlmsapi/routes/region"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	mainRouter := app.Group("/api/v2", logger.New())
	auth.SetupAuthRoute(&mainRouter)
	files.SetupFilesRoute(&mainRouter)
	level.SetupLevelRoute(mainRouter)
	course.SetupCourseRouter(&mainRouter)
	city.SetupCityRoutes(mainRouter)
	region.SetupRegionRoutes(&mainRouter)
}
