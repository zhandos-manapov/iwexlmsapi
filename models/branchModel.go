package models

type BranchDB struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	AddressID int    `json:"address_id" db:"address_id"`
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
