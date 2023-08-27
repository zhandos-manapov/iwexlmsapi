package main

import (
	"iwexlmsapi/routes/auth"
	"iwexlmsapi/routes/course"
	"iwexlmsapi/routes/lesson"
	"iwexlmsapi/routes/level"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	mainRouter := app.Group("/api/v2", logger.New())
	auth.SetupAuthRoute(mainRouter)
	level.SetupLevelRoute(mainRouter)
	course.SetupCourseRouter(mainRouter)
	lesson.SetupLessonRouter(mainRouter)
}
