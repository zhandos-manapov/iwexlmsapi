package models

import "github.com/jackc/pgx/v5/pgtype"

type BranchDB struct {
	ID        pgtype.Numeric `json:"id" db:"id"`
	Name      pgtype.Text    `json:"name" db:"name"`
	AddressID pgtype.Numeric `json:"address_id" db:"address_id"`
}

type CreateBranchDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required,min=2,max=50"`
	AddressID int    `json:"address_id"`
}

type UpdateBranchDTO struct {
	Name      string `json:"name" validate:"omitempty,min=2,max=50"`
	AddressID int    `json:"address_id"`
}
