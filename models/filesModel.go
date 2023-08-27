package models

type FileOperationsReqBody struct {
	Action          string `json:"action"`
	Path            string `json:"path"`
	ShowHiddenItems bool   `json:"showHiddenItems"`
}


