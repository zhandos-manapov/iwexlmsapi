package auth

import (
	"github.com/gofiber/fiber/v2"
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"
)

func SetupAuthRouter(router *fiber.Router) {
	authRouter := (*router).Group("/auth")

	authRouter.Post("/signin", middleware.BodyParserValidatorMiddleware(&models.UserSignInDTO{}), signIn)
	authRouter.Post("/signup", middleware.BodyParserValidatorMiddleware(&models.UserSignUpDTO{}), signUp)
}
