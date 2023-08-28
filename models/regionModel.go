package models

type Region struct {
	ID         int    `json:"id" db:"id"`
	RegionName string `json:"region_name" db:"region_name"`
	CountyID   int    `json:"country_id" db:"country_id"`
}
