package main

import (
	"iwexlmsapi/routes/auth"
	"iwexlmsapi/routes/city"
	"iwexlmsapi/routes/country"
	"iwexlmsapi/routes/course"
	"iwexlmsapi/routes/files"
	"iwexlmsapi/routes/level"
	"iwexlmsapi/routes/region"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	mainRouter := app.Group("/api/v2", logger.New())
	auth.SetupAuthRouter(&mainRouter)
	files.SetupFilesRouter(&mainRouter)
	level.SetupLevelRouter(&mainRouter)
	course.SetupCourseRouter(&mainRouter)
	city.SetupCityRouter(&mainRouter)
	region.SetupRegionRouter(&mainRouter)
	country.SetupCountryRouter(&mainRouter)
}
