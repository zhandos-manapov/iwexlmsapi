package users

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"
	users "iwexlmsapi/routes/users/toggle"

	"github.com/gofiber/fiber/v2"
)

func SetupUsersRoutes(router *fiber.Router) {
	usersRouter := (*router).Group("/user")

	usersRouter.Get("/", FindMany)
	usersRouter.Get("/:id", FindOne)

	usersRouter.Post("/", middleware.BodyParserValidatorMiddleware(&models.User{}), CreateOne)
	usersRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.User{}), UpdateOne)
	usersRouter.Delete("/:id", DeleteOne)

	usersRouter.Post("/:id/toggle", users.Toggle)
}
