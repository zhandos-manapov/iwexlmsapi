package auth

import (
	"github.com/gofiber/fiber/v2"
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"
)

func SetupAuthRoute(router fiber.Router) {
	authRouter := router.Group("/auth")

	authRouter.Post("/signin", middleware.BodyParserValidatorMiddleware(&models.UserLog{}), signIn)
	authRouter.Post("/signup", middleware.BodyParserValidatorMiddleware(&models.User{}), signUp)
}
