package auth

import (
	"iwexlmsapi/middleware"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRouter(router *fiber.Router) {
	authRouter := (*router).Group("/auth")

	authRouter.Post("/signin", middleware.BodyParserValidatorMiddlewareForStruct(&models.UserSignInDTO{}), signIn)
	authRouter.Post("/signup", middleware.BodyParserValidatorMiddlewareForStruct(&models.UserSignUpDTO{}), signUp)
}
