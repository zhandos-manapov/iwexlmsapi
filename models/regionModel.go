package models

import "github.com/jackc/pgx/v5/pgtype"

type CreateRegionDTO struct {
	ID         int    `json:"id"`
	RegionName string `json:"region_name" validate:"required"`
	CountyID   int    `json:"country_id" validate:"required"`
}

type UpdateRegionDTO struct {
	RegionName string `json:"region_name"`
	CountyID   int    `json:"country_id"`
}

type RegionDB struct {
	ID         pgtype.Numeric `json:"id" db:"id"`
	RegionName pgtype.Text    `json:"region_name" db:"region_name"`
	CountyID   pgtype.Numeric `json:"country_id" db:"country_id"`
}
