package country

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
)

func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "SELECT * FROM country WHERE id=$1"
	country := models.CountryDB{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(&country.ID, &country.CountryName); err != nil {
		return err
	}
	return c.JSON(country)
}

func findMany(c *fiber.Ctx) error {
	query := "SELECT * FROM country"
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	countries, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.CountryDB])
	if len(countries) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Страны не найдены")
	}
	return c.JSON(countries)
}

func createOne(c *fiber.Ctx) error {
	country := c.Locals("body").(*models.CreateCountryDTO)
	query := "INSERT INTO country (country_name) VAlUES ($1) RETURNING id"
	if err := database.Pool.QueryRow(context.Background(), query, country.CountryName).Scan(&country.ID); err != nil {
		return err
	}
	return c.JSON(country)
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	country := c.Locals("body").(*models.UpdateCountryDTO)
	if country.CountryName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления")
	}
	query := "UPDATE country SET country_name=$1 WHERE id=$2"
	if tag, err := database.Pool.Exec(context.Background(), query, country.CountryName, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Страна успешно обновлена"})
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM country WHERE id=$1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Страна успешно удалена"})
}
