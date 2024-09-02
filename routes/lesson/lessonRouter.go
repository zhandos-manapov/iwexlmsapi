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
	localRouter.Post("/", middleware.BodyParserValidatorMiddlewareForSlice([]models.CreateLessonDTO{}), createMany)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddlewareForStruct(&models.UpdateLessonDTO{}), updateOne)
	localRouter.Delete("/:id", deleteOne)
}
