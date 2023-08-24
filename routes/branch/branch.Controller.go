package branch

import (
	"context"
	"database/sql"
	"iwexlmsapi/database"
	"iwexlmsapi/errors"
	"iwexlmsapi/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func findMany(c *fiber.Ctx) error {
    query := "SELECT * FROM branch_office"
    rows, err := database.Pool.Query(context.Background(), query)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(errors.ServerError{Message: err.Error()})
    }
    defer rows.Close()

    branchOffices := []models.BranchOffice{}
    for rows.Next() {
        branchOffice := models.BranchOffice{}
        var id int
        var name sql.NullString
        var address sql.NullString
        if err := rows.Scan(&id, &name, &address); err != nil {
            return c.Status(http.StatusInternalServerError).JSON(errors.ServerError{Message: err.Error()})
        }
        branchOffice.ID = id
        if name.Valid {
            branchOffice.Name = name.String
        }
        if address.Valid {
            branchOffice.Address = address.String
        }
        branchOffices = append(branchOffices, branchOffice)
    }

    if len(branchOffices) == 0 {
        return c.Status(http.StatusNotFound).JSON(errors.NotFoundError{Message: "Филиалы не найдены"})
    }

    return c.JSON(branchOffices)
}


func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "SELECT * FROM branch_office WHERE id = $1"
	row := database.Pool.QueryRow(context.Background(), query, id)
	branchOffice :=  models.BranchOffice{}
	err := row.Scan(&branchOffice.ID, &branchOffice.Name, &branchOffice.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(errors.NotFoundError{Message: "Филиал с id " + id + " не найден"})
		}
		return c.Status(http.StatusInternalServerError).JSON(errors.ServerError{Message: err.Error()})
	}

	return c.JSON(branchOffice)
}

func createOne(c *fiber.Ctx) error {
	branchOffice := new(models.BranchOffice)
	if err := c.BodyParser(branchOffice); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.BadRequestError{Message: err.Error()})
	}

	query := "INSERT INTO branch_office (name, address_id) VALUES ($1, $2) RETURNING id"
	row := database.Pool.QueryRow(context.Background(), query, branchOffice.Name, branchOffice.Address)
	if err := row.Scan(&branchOffice.ID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.ServerError{Message: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(branchOffice)
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM branch_office WHERE id = $1"
	_, err := database.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.ServerError{Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Успешно удалено"})
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	branchOffice := new(models.BranchOffice)
	if err := c.BodyParser(branchOffice); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.BadRequestError{Message: err.Error()})
	}

	query := "UPDATE branch_office SET name = $1 WHERE id = $2"
	_, err := database.Pool.Exec(context.Background(), query, branchOffice.Name, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.ServerError{Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Успешно обновлено"})
}
