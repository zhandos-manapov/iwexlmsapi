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
	row := database.Pool.QueryRow(context.Background(), query, id)

	var city models.City
	if err := row.Scan(&city.ID, &city.CityName, &city.RegionID); err != nil {
		if err == pgx.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Город не найден",
			})
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(city)
}

func FindAll(c *fiber.Ctx) error {
	query := "SELECT * FROM city"
	rows, err := database.Pool.Query(context.Background(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer rows.Close()

	cities := make([]models.City, 0)
	for rows.Next() {
		var city models.City
		if err := rows.Scan(&city.ID, &city.CityName, &city.RegionID); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		cities = append(cities, city)
	}

	if len(cities) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Города не найдены",
		})
	}

	return c.Status(fiber.StatusOK).JSON(cities)
}

func CreateOne(c *fiber.Ctx) error {
	var city models.City
	if err := c.BodyParser(&city); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка разбора данных",
		})
	}

	query := "INSERT INTO city (city_name, region_id) VALUES($1, $2)"
	_, err := database.Pool.Exec(context.Background(), query, city.CityName, city.RegionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Данное название уже существует",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Город успешно добавлен"})
}

func UpdateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var city models.City
	if err := c.BodyParser(&city); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка разбора данных",
		})
	}

	query := "UPDATE city SET city_name=$1 WHERE id=$2"
	_, err := database.Pool.Exec(context.Background(), query, city.CityName, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка в изменениях данных",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Успешно обновлено"})
}

func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM city WHERE id=$1"
	_, err := database.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка при удалении города",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Город успешно удален"})
}
