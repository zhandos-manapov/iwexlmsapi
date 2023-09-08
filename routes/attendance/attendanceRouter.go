package attendance

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupAttendanceRouter(router *fiber.Router) {
	localRouter := (*router).Group("/attendance")
	localRouter.Get("/:id", findOne)
	localRouter.Get("/", findMany)
	localRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.Attendance{}), createOne)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdAttendance{}), updateOne)
	localRouter.Delete("/:id", deleteOne)
}
