package models

type BranchOffice struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}