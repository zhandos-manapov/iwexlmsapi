package class

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupClassRouter(router *fiber.Router) {
	classRouter := (*router).Group("/classes")
	classRouter.Get("/", findMany)
	classRouter.Get("/:id", findOne)
	classRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.CreateClassDTO{}), createOne)
	classRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateClassDTO{}), updateOne)
	classRouter.Delete("/:id", deleteOne)
	classRouter.Get("/:id/people", getEnrolledStudents)
	classRouter.Post("/:id/people", middleware.BodyParserValidatorMiddleware(&models.EnrollStudentsDTO{}), enrollStudents)
}
