package users

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
)

func FindOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `
	SELECT users.id,
		users.first_name,
		users.last_name,
		users.email,
		users.contact_number,
		users.date_of_birth,
		users.is_active,
		users.role,
		role.role_name
	FROM users
		INNER JOIN role ON users.role = role.id
	WHERE (users.id = $1)`
	user := models.User{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.ContactNumber,
		&user.DateOfBirth,
		&user.IsActive,
		&user.Role,
		&user.RoleName,
	); err != nil {
		return err
	}
	return c.JSON(user)
}

func FindMany(c *fiber.Ctx) error {
	query := `
	SELECT users.id,
		users.first_name,
		users.last_name,
		users.email,
		users.contact_number,
		users.date_of_birth,
		users.is_active,
		users.role,
		role.role_name
	FROM users
		INNER JOIN role ON users.role = role.id`
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.User])
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Пользователи не найдены ")
	}
	return c.JSON(users)
}

func UpdateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := c.Locals("body", &user); err != nil {
		return err.(error)
	}

	query := `
        UPDATE users SET first_name=$1, last_name=$2, email=$3, contact_number=$4, date_of_birth=$5, is_active=$6, role_name=$7
        WHERE id=$8`
	_, err := database.Pool.Exec(context.Background(), query,
		user.FirstName, user.LastName, user.Email, user.ContactNumber, user.DateOfBirth, user.IsActive, user.RoleName, id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "User successfully updated"})
}

func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `DELETE FROM users WHERE id = $1`
	_, err := database.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "User successfully deleted"})
}
