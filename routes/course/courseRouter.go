package course

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupCourseRouter(router *fiber.Router) {
	localRouter := (*router).Group("/courses")
	localRouter.Get("/:id", FindOne)
	localRouter.Get("/", FindMany)
	localRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateCourse{}), CreateOne)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateCourse{}), UpdateOne)
	localRouter.Delete("/:id", DeleteOne)
}
