package middleware

import (
	"fmt"
	"iwexlmsapi/models"
	"iwexlmsapi/xvalidator"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const reqBody = "body"

type reqBodyType interface {
	models.User |
		models.UserLog |
		models.Level |
		models.CourseCreate |
		models.CourseUpdate |
		models.FileOperationsReqBody |
		models.City |
		models.Region |
		models.Country |
		models.CreateLesson
}

func BodyParserValidatorMiddleware[T reqBodyType](data *T) func(c *fiber.Ctx) error {
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
