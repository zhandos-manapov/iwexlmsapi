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
	localRouter.Post("/", middleware.BodyParserValidatorMiddlewareForStruct(&models.CreateCourseDTO{}), createOne)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddlewareForStruct(&models.UpdateCourseDTO{}), updateOne)
	localRouter.Delete("/:id", deleteOne)
	localRouter.Get("/:id/classes", findClassesByCourse)
}
