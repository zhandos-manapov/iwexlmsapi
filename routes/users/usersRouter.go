package users

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(router *fiber.Router) {
	(*router).Get("/filter", filterUsers)

	usersRouter := (*router).Group("/users")

	usersRouter.Get("/", findMany)
	usersRouter.Get("/:id", findOne)
	usersRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UpdateUserDTO{}), updateOne)
	usersRouter.Delete("/:id", deleteOne)
}
