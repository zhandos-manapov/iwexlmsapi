package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"iwexlmsapi/utils"
	"regexp"
	"strings"
)

func Authorize(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	token := headers["Authorization"]
	match, err := regexp.MatchString(`Bearer \S+\.\S+\.\S+`, token)
	if err != nil {
		return err
	}
	if !match {
		return fiber.NewError(fiber.StatusUnauthorized, "Недействительный токен")
	}
	tokenParts := strings.Split(token, " ")
	tkn, err := jwt.Parse(tokenParts[1], func(t *jwt.Token) (interface{}, error) {
		return utils.PrivateKey.Public(), nil
	})
	if err != nil {
		return err
	}
	c.Locals("user", tkn.Claims)
	return c.Next()
}
