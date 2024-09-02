package main

import (
	"iwexlmsapi/routes/auth"
	"iwexlmsapi/routes/branch"
	"iwexlmsapi/routes/city"
	"iwexlmsapi/routes/class"
	"iwexlmsapi/routes/country"
	"iwexlmsapi/routes/course"
	"iwexlmsapi/routes/files"
	"iwexlmsapi/routes/lesson"
	"iwexlmsapi/routes/level"
	"iwexlmsapi/routes/quiz"
	"iwexlmsapi/routes/region"
	"iwexlmsapi/routes/users"

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
	lesson.SetupLessonRouter(&mainRouter)
	class.SetupClassRouter(&mainRouter)
	users.SetupUserRouter(&mainRouter)
	branch.SetupBranchRouter(&mainRouter)
	class.SetupClassRouter(&mainRouter)
	quiz.SetupQuizRouter(&mainRouter)
}
