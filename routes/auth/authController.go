package auth

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"iwexlmsapi/utils"
)

func signIn(c *fiber.Ctx) error {
	body := c.Locals("body")
	user := body.(*models.UserSignInDTO)
	query := `
	SELECT users.id,
		users.email,
		users.hash,
		users.salt,
		role.role_name,
		users.is_active
	FROM users
		INNER JOIN role ON users.role = role.id
	WHERE users.email = $1`
	dbUser := models.UserDB{}
	if err := database.Pool.QueryRow(context.Background(), query, user.Email).Scan(&dbUser.Id, &dbUser.Email, &dbUser.Hash, &dbUser.Salt, &dbUser.RoleName, &dbUser.IsActive); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Пользователь не найден")
	}
	if dbUser.IsActive.Bool == false {
		return fiber.NewError(fiber.StatusForbidden, "Дождитесь одобрения администратора")
	}
	validPass := utils.ValidPassword(user.Password, dbUser.Hash.String, dbUser.Salt.String)
	if !validPass {
		return fiber.NewError(fiber.StatusUnauthorized, "Неверный пароль")
	}
	tokenStruct, err := utils.IssueJWT(&dbUser)
	if err != nil {
		return err
	}
	return c.JSON(tokenStruct)
}

func signUp(c *fiber.Ctx) error {
	body := c.Locals("body")
	user := body.(models.UserSignUpDTO)
	query := "SELECT email FROM users WHERE email = $1"
	result, err := database.Pool.Query(context.Background(), query, user.Email)
	if err != nil {
		return err
	}
	users, err := pgx.CollectRows[string](result, pgx.RowTo[string])
	if err != nil {
		return err
	}
	if len(users) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Email уже существует")
	}
	hash, salt, err := utils.GenHash(user.Password)
	if err != nil {
		return err
	}
	query = `
		INSERT INTO users (first_name, last_name, email, contact_number, date_of_birth, salt, hash) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	commandTag, err := database.Pool.Exec(context.Background(), query, user.FirstName, user.LastName, user.Email, user.ContactNumber, user.DateOfBirth, salt, hash)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Вы успешно зарегистрированы"})
}
