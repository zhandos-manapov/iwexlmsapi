package users

import (
	"github.com/gofiber/fiber/v2"
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"
)

func SetupUserRouter(router *fiber.Router) {
	usersRouter := (*router).Group("/users")

	usersRouter.Get("/", findMany)
	usersRouter.Get("/:id", findOne)
	usersRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateUserDTO{}), updateOne)
	usersRouter.Delete("/:id", deleteOne)
}
