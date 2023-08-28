package city

import (
	"context"
	"iwexlmsapi/database"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func FindOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "SELECT * FROM city WHERE id=$1"
	city := models.City{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(&city.ID, &city.CityName, &city.RegionID); err != nil {
		return err
	}
	return c.JSON(city)
}

func FindMany(c *fiber.Ctx) error {
	query := "SELECT * FROM city"
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	cities, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.City])
	if err != nil {
		return err
	}
	if len(cities) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Города не найдены")
	}
	return c.JSON(cities)
}

func CreateOne(c *fiber.Ctx) error {
	city := c.Locals("body").(*models.City)
	query := "INSERT INTO city (city_name, region_id) VALUES($1, $2)"
	if tag, err := database.Pool.Exec(context.Background(), query, city.CityName, city.RegionID); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Город успешно добавлен"})
}

func UpdateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	city := c.Locals("body").(*models.City)
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

func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM city WHERE id=$1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Город успешно удален"})
}
