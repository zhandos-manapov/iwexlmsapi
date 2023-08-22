package models

type City struct {
	ID       int    `json:"id"`
	CityName string `json:"city_name" validate:"required,min=2,max=50"`
	RegionID int    `json:"region_id" validate:"required"`
}
