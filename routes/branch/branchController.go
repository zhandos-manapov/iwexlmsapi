package branch

import (
	"context"
	"fmt"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

func findMany(c *fiber.Ctx) error {
	query := `
	SELECT id,
		name,
		COALESCE(address_id, 0) as address_id
	from branch_office`
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	branches, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.BranchDB])
	if len(branches) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Филиалы не найдены ")
	}
	return c.JSON(branches)
}

func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `
	SELECT id,
		name,
		COALESCE(address_id, 0) as address_id
	FROM branch_office 
	WHERE id = $1`
	branch := models.BranchDB{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(&branch.ID, &branch.Name, &branch.AddressID); err != nil {
		return err
	}
	return c.JSON(branch)
}

func createOne(c *fiber.Ctx) error {
	branch := c.Locals("body").(*models.CreateBranchDTO)
	query := "INSERT INTO branch_office (name, address_id) VALUES ($1, $2) RETURNING id"
	if err := database.Pool.QueryRow(context.Background(), query, branch.Name, zeronull.Int8(branch.AddressID)).Scan(&branch.ID); err != nil {
		return err
	}
	return c.JSON(branch)
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM branch_office WHERE id = $1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Филиал не найден")
	}
	return c.JSON(models.RespMsg{Message: "Филиал успешно удален"})
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	branch := c.Locals("body").(*models.UpdateBranchDTO)
	if (*branch == models.UpdateBranchDTO{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления")
	}
	query := strings.Builder{}
	query.WriteString("UPDATE branch_office SET")
	queryParams := []any{id}

	if branch.Name != "" {
		query.WriteString(fmt.Sprintf(" name=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, branch.Name)
	}

	if branch.AddressID != 0 {
		query.WriteString(fmt.Sprintf(" address_id=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, branch.AddressID)
	}

	queryString := query.String()
	queryString = queryString[:len(queryString)-1]
	queryString += " WHERE id = $1"

	if tag, err := database.Pool.Exec(context.Background(), queryString, queryParams...); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Филиал не найден")
	}
	return c.JSON(models.RespMsg{Message: "Филиал успешно обновлен"})
}
