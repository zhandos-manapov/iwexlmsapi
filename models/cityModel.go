package models

import "github.com/jackc/pgx/v5/pgtype"

type CityDB struct {
	ID       pgtype.Numeric `json:"id" db:"id"`
	CityName pgtype.Text    `json:"city_name" db:"city_name"`
	RegionID pgtype.Numeric `json:"region_id" db:"region_id"`
}

type CreateCityDTO struct {
	ID       int    `json:"id"`
	CityName string `json:"city_name" validate:"required,min=2,max=50"`
	RegionID string `json:"region_id" validate:"required"`
}

type UpdateCityDTO struct {
	CityName string `json:"city_name" validate:"omitempty,min=2,max=50"`
	RegionID string `json:"region_id"`
}
