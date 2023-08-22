package auth

import (
	"context"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"iwexlmsapi/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func signIn(c *fiber.Ctx) error {
	body := c.Locals("body")
	user := body.(*models.UserLog)
	query := `
    SELECT users.id, users.email, users.hash, users.salt, role.role_name, users.is_active
    FROM users
    INNER JOIN role ON users.role = role.id
    WHERE users.email=$1`
	dbUser := models.User{}
	if err := database.Pool.QueryRow(context.Background(), query, user.Email).Scan(&dbUser.Id, &dbUser.Email, &dbUser.Hash, &dbUser.Salt, &dbUser.RoleName, &dbUser.IsActive); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Пользователь не найден")
	}
	if dbUser.IsActive == false {
		return fiber.NewError(fiber.StatusForbidden, "Дождитесь одобрения администратора")
	}
	validPass := utils.ValidPassword(user.Password, dbUser.Hash, dbUser.Salt)
	if !validPass {
		return fiber.NewError(fiber.StatusUnauthorized, "Неверный пароль")
	}
	//issue jwt

	return c.SendString("auth route is working")
}

func signUp(c *fiber.Ctx) error {
	body := c.Locals("body")
	user := body.(*models.User)
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
	return c.JSON(models.ServerError{Message: "Вы успешно зарегистрированы"})
}
