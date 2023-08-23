package users

import (
	"context"
	"errors"
	"iwexlmsapi/database"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func FindMany(c *fiber.Ctx) error {
	query := `
		SELECT users.id, users.first_name, users.last_name, users.email, users.contact_number, users.date_of_birth, users.is_active, users.role_name
		FROM users`
	rows, err := database.Pool.Query(context.Background(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.Id, &user.FirstName, &user.LastName, &user.Email,
			&user.ContactNumber, &user.DateOfBirth, &user.IsActive, &user.RoleName,
		); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		users = append(users, user)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func FindOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `
		SELECT users.id, users.first_name, users.last_name, users.email, users.contact_number, users.date_of_birth, users.is_active, users.role_name
		FROM users
		WHERE (users.id=$1)`
	row := database.Pool.QueryRow(context.Background(), query, id)

	var user models.User
	if err := row.Scan(
		&user.Id, &user.FirstName, &user.LastName, &user.Email,
		&user.ContactNumber, &user.DateOfBirth, &user.IsActive, &user.RoleName,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func CreateOne(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing data",
		})
	}

	emailExistsQuery := `SELECT email FROM users WHERE email = $1`
	emailCheckRow := database.Pool.QueryRow(context.Background(), emailExistsQuery, user.Email)
	var existingEmail string
	if err := emailCheckRow.Scan(&existingEmail); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}

	query := `
        INSERT INTO users (first_name, last_name, email, contact_number, date_of_birth, is_active, role_name)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := database.Pool.Exec(context.Background(), query,
		user.FirstName, user.LastName, user.Email, user.ContactNumber, user.DateOfBirth, user.IsActive, user.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User successfully created"})
}

func UpdateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing data",
		})
	}

	query := `
        UPDATE users SET first_name=$1, last_name=$2, email=$3, contact_number=$4, date_of_birth=$5, is_active=$6, role_name=$7
        WHERE id=$8`
	_, err := database.Pool.Exec(context.Background(), query,
		user.FirstName, user.LastName, user.Email, user.ContactNumber, user.DateOfBirth, user.IsActive, user.RoleName, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User successfully updated"})
}

func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `DELETE FROM users WHERE id = $1`
	_, err := database.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User successfully deleted"})
}
