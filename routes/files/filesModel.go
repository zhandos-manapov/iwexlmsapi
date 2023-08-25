package files

import "time"

type fileOperationsReqBody struct {
	Action          string `json:"action"`
	Path            string `json:"path"`
	ShowHiddenItems bool   `json:"showHiddenItems"`
}

type fileStruct struct {
	Name         string    `json:"name"`
	Size         string    `json:"size"`
	IsFile       bool      `json:"isFile"`
	DateModified time.Time `json:"dateModified"`
	DateCreated  time.Time `json:"dateCreated"`
	Type         string    `json:"type"`
	FilterPath   string    `json:"filterPath"`
	Permission   any       `json:"permission"`
	HasChild     bool      `json:"hasChild"`
}
