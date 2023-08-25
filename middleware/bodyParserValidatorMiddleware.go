package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"iwexlmsapi/xvalidator"
	"strings"
)

const reqBody = "body"

func BodyParserValidatorMiddleware[T interface{}](data *T) func(c *fiber.Ctx) error {
func BodyParserValidatorMiddleware[T models.User | models.UserLog | models.Level | models.CourseSend](data *T) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(data); err != nil {
			return err
		}
		if errs := xvalidator.ValidateStruct(data); len(errs) > 0 && errs[0].Error {
			ln := len(errs)
			errMessages := strings.Builder{}
			for i := 0; i < ln; i++ {
				err := errs[i]
				str := fmt.Sprintf(
					"[%s]: '%v' | Needs to implement '%s'\n",
					err.FailedField,
					err.Value,
					err.Tag,
				)
				errMessages.WriteString(str)
			}
			return fiber.NewError(fiber.StatusBadRequest, errMessages.String())
		}
		c.Locals(reqBody, data)
		return c.Next()
	}
}
