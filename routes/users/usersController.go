package users

import (
	"context"
	"fmt"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func findOne(c *fiber.Ctx) error {
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
	user := models.UserDB{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&user.Id,
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

func findMany(c *fiber.Ctx) error {
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
	users, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[models.UserDB])
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Пользователи не найдены ")
	}
	return c.JSON(users)
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	user := c.Locals("body").(*models.UpdateUserDTO)

	if (*user == models.UpdateUserDTO{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления")
	}

	isActive := false
	q := "SELECT is_active FROM users WHERE id = $1"
	if err := database.Pool.QueryRow(context.Background(), q, id).Scan(&isActive); err != nil {
		return err
	}

	query := strings.Builder{}
	query.WriteString("UPDATE users SET")
	queryParams := []any{id}

	if user.FirstName != "" {
		query.WriteString(fmt.Sprintf(" first_name=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.FirstName)
	}

	if user.LastName != "" {
		query.WriteString(fmt.Sprintf(" last_name=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.LastName)
	}

	if user.ContactNumber != "" {
		query.WriteString(fmt.Sprintf(" contact_number=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.ContactNumber)
	}

	if user.DateOfBirth != "" {
		query.WriteString(fmt.Sprintf(" date_of_birth=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.DateOfBirth)
	}

	if user.Email != "" {
		query.WriteString(fmt.Sprintf(" email=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.Email)
	}

	if user.Role != 0 {
		query.WriteString(fmt.Sprintf(" role=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.Role)
	}

	if user.IsActive != isActive {
		query.WriteString(fmt.Sprintf(" is_active=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.IsActive)
	}

	queryString := query.String()
	queryString = queryString[:len(queryString)-1]
	queryString += " WHERE id=$1"

	if tag, err := database.Pool.Exec(context.Background(), queryString, queryParams...); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Пользователь не найден")
	}
	return c.JSON(models.RespMsg{Message: "Данные пользователя успешно обновлены"})
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `DELETE FROM users WHERE id = $1`
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Пользователь не найден")
	}
	return c.JSON(models.RespMsg{Message: "Пользователь успешно удален "})
}
