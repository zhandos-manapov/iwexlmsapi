package level

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
)

func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "SELECT * FROM level WHERE id=$1"
	level := models.Level{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(&level.ID, &level.LevelName); err != nil {
		return err
	}
	return c.JSON(level)
}

func findMany(c *fiber.Ctx) error {
	query := "SELECT * FROM level"
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	levels, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Level])
	if err != nil {
		return err
	}
	if len(levels) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Уровни не найдены")
	}
	return c.JSON(levels)
}

func createOne(c *fiber.Ctx) error {
	level := c.Locals("body").(*models.Level)
	query := "INSERT INTO level (level_name) VALUES($1)"
	if tag, err := database.Pool.Exec(context.Background(), query, level.LevelName); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Уровень успешно создан"})
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	level := c.Locals("body").(*models.Level)
	if level.LevelName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления")
	}
	query := "UPDATE level SET level_name=$1 WHERE id=$2"
	if tag, err := database.Pool.Exec(context.Background(), query, level.LevelName, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Уровень успешно обновлен"})
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM level WHERE id=$1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Уровень успешно удален"})
}
