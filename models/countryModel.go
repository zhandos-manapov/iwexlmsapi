package models

type CreateCountryDTO struct {
	ID          int    `json:"id"`
	CountryName string `json:"country_name" validate:"required,min=2"`
}

type UpdateCountryDTO struct {
	CountryName string `json:"country_name"`
}

type CountryDB struct {
	ID          int    `json:"id" db:"id"`
	CountryName string `json:"country_name" db:"country_name"`
}
