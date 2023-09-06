package lesson

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupLessonRouter(router *fiber.Router) {
	localRouter := (*router).Group("/lessons")
	localRouter.Get("/gil/:id", getIdLesson)
	localRouter.Get("/:id", findOne)
	localRouter.Get("/", findMany)
	localRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateLessonDTO{}), createOne)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateLessonDTO{}), updateOne)
	localRouter.Delete("/:id", deleteOne)
}
