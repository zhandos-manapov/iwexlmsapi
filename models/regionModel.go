package models

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
	ID         int    `json:"id" db:"id"`
	RegionName string `json:"region_name" db:"region_name"`
	CountyID   int    `json:"country_id" db:"country_id"`
}
