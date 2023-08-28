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
	region := models.Region{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(&region.ID, &region.RegionName, &region.CountyID); err != nil {
		return err
	}
	return c.JSON(region)
}

func FindMany(c *fiber.Ctx) error {
	query := "SELECT * FROM region"
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	regions, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Region])
	if err != nil {
		return err
	}
	if len(regions) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Регионы не найдены")
	}
	return c.JSON(regions)
}

func CreateOne(c *fiber.Ctx) error {
	region := c.Locals("body").(*models.Region)
	query := "INSERT INTO region (region_name, county_id) VALUES($1, $2)"
	if tag, err := database.Pool.Exec(context.Background(), query, region.RegionName, region.CountyID); err != nil {

	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Регион успешно добавлен"})
}

func UpdateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	region := c.Locals("body").(*models.Region)
	if region.RegionName == "" && region.CountyID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления")
	}
	query := "UPDATE region SET region_name=$1, county_id=$2 WHERE id=$3"
	if tag, err := database.Pool.Exec(context.Background(), query, region.RegionName, region.CountyID, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Регион успешно обновлен"})
}

func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM region WHERE id=$1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Регион успешно удален"})
}
