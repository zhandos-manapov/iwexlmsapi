package attendance

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupAttendanceRouter(router *fiber.Router) {
	localRouter := (*router).Group("/attendances")
	localRouter.Get("/:id", findOne)
	localRouter.Get("/:id/class", findMany)
	localRouter.Post("/", middleware.BodyParserValidatorMiddlewareForSlice([]models.UpdAttendance{}), createMany)
	localRouter.Patch("/:id", middleware.BodyParserValidatorMiddlewareForStruct(&models.UpdAttendance{}), updateOne)
	localRouter.Delete("/:id", deleteOne)
}
