package models

type Country struct {
	ID          string `json:"id" db:"id"`
	CountryName string `json:"country_name" db:"country_name"`
}
