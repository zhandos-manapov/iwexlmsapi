package city

import (
	"context"
	"iwexlmsapi/database"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "SELECT * FROM city WHERE id=$1"
	city := models.CityDB{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(&city.ID, &city.CityName, &city.RegionID); err != nil {
		return err
	}
	return c.JSON(city)
}

func findMany(c *fiber.Ctx) error {
	query := "SELECT * FROM city"
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	cities, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.CityDB])
	if err != nil {
		return err
	}
	if len(cities) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Города не найдены")
	}
	return c.JSON(cities)
}

func createOne(c *fiber.Ctx) error {
	city := c.Locals("body").(*models.CreateCityDTO)
	query := "INSERT INTO city (city_name, region_id) VALUES($1, $2) RETURNING id"
	if err := database.Pool.QueryRow(context.Background(), query, city.CityName, city.RegionID).Scan(&city.ID); err != nil {
		return err
	}
	return c.JSON(city)
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	city := c.Locals("body").(*models.UpdateCityDTO)
	if city.CityName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления")
	}
	query := "UPDATE city SET city_name=$1 WHERE id=$2"
	if tag, err := database.Pool.Exec(context.Background(), query, city.CityName, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Город успешно обновлен"})
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM city WHERE id=$1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Город успешно удален"})
}
