package lesson

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupLessonRouter(router fiber.Router) {
	localRouter := router.Group("/lesson")
	localRouter.Get("/:id", findOne)
	localRouter.Get("/", findMany)
	localRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateLesson{}), CreateOne)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.CreateLesson{}), updateOne)
	localRouter.Delete("/:id", deleteOne)
}
