package users

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"
	users "iwexlmsapi/routes/users/toggle"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(router *fiber.Router) {
	usersRouter := (*router).Group("/user")

	usersRouter.Get("/", FindMany)
	usersRouter.Get("/:id", FindOne)
	usersRouter.Patch("/:id", middleware.BodyParserValidatorMiddleware(&models.UserUpdate{}), UpdateOne)
	usersRouter.Delete("/:id", DeleteOne)
	usersRouter.Post("/:id/toggle", users.Toggle)
}
