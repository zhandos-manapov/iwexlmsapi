package models

type Region struct {
	ID         int    `json:"id"`
	RegionName string `json:"region_name"`
	CountyID   int    `json:"county_id"`
}
