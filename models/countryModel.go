package models

import "github.com/jackc/pgx/v5/pgtype"

type CreateCountryDTO struct {
	ID          int    `json:"id"`
	CountryName string `json:"country_name" validate:"required,min=2"`
}

type UpdateCountryDTO struct {
	CountryName string `json:"country_name"`
}

type CountryDB struct {
	ID          pgtype.Numeric `json:"id" db:"id"`
	CountryName pgtype.Text    `json:"country_name" db:"country_name"`
}
