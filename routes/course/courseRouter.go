package course

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupCourseRouter(router *fiber.Router) {
	localRouter := (*router).Group("/courses")
	localRouter.Get("/:id", findOne)
	localRouter.Get("/", findMany)
	localRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CourseCreate{}), createOne)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.CourseUpdate{}), updateOne)
	localRouter.Delete("/:id", deleteOne)
}
