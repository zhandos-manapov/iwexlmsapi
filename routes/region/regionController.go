package region

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
)

func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "SELECT * FROM region WHERE id=$1"
	region := models.RegionDB{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(&region.ID, &region.RegionName, &region.CountyID); err != nil {
		return err
	}
	return c.JSON(region)
}

func findMany(c *fiber.Ctx) error {
	query := "SELECT * FROM region"
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	regions, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.RegionDB])
	if err != nil {
		return err
	}
	if len(regions) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Регионы не найдены")
	}
	return c.JSON(regions)
}

func createOne(c *fiber.Ctx) error {
	region := c.Locals("body").(*models.CreateRegionDTO)
	query := "INSERT INTO region (region_name, country_id) VALUES($1, $2)"
	if tag, err := database.Pool.Exec(context.Background(), query, region.RegionName, region.CountyID); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Регион успешно добавлен"})
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	region := c.Locals("body").(*models.UpdateRegionDTO)
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

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM region WHERE id=$1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Регион успешно удален"})
}
