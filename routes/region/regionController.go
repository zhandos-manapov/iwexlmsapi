package region

import (
	"context"
	"iwexlmsapi/database"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func FindOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "SELECT * FROM region WHERE id=$1"
	row := database.Pool.QueryRow(context.Background(), query, id)

	var region models.Region
	if err := row.Scan(&region.ID, &region.RegionName, &region.CountyID); err != nil {
		if err == pgx.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Регион не найден",
			})
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(region)
}

func FindAll(c *fiber.Ctx) error {
	query := "SELECT * FROM region"
	rows, err := database.Pool.Query(context.Background(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer rows.Close()

	regions := make([]models.Region, 0)
	for rows.Next() {
		var region models.Region
		if err := rows.Scan(&region.ID, &region.RegionName, &region.CountyID); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		regions = append(regions, region)
	}

	if len(regions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Регионы не найдены",
		})
	}

	return c.Status(fiber.StatusOK).JSON(regions)
}

func CreateOne(c *fiber.Ctx) error {
	var region models.Region
	if err := c.BodyParser(&region); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка разбора данных",
		})
	}

	query := "INSERT INTO region (region_name, county_id) VALUES($1, $2)"
	_, err := database.Pool.Exec(context.Background(), query, region.RegionName, region.CountyID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Данное название уже существует",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Регион успешно добавлен"})
}

func UpdateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var region models.Region
	if err := c.BodyParser(&region); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка разбора данных",
		})
	}

	query := "UPDATE region SET region_name=$1, county_id=$2 WHERE id=$3"
	_, err := database.Pool.Exec(context.Background(), query, region.RegionName, region.CountyID, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка в изменениях данных",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Успешно обновлено"})
}

func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM region WHERE id=$1"
	_, err := database.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка при удалении региона",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Регион успешно удален"})
}
