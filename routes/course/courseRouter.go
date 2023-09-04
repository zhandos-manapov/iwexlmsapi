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
	localRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateCourseDTO{}), createOne)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateCourseDTO{}), updateOne)
	localRouter.Delete("/:id", deleteOne)
}
