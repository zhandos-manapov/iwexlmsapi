package course

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupCourseRouter(router fiber.Router) {
	localRouter := router.Group("/course")
	localRouter.Get("/:id", FindOne)
	localRouter.Get("/", FindMany)
	localRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CourseSend{}), CreateOne)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.CourseSend{}), UpdateOne)
	localRouter.Delete("/:id", DeleteOne)
}
