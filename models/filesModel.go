package models

import "time"

type FileOperationsReqBody struct {
	Action          string        `json:"action"`
	Path            string        `json:"path"`
	ShowHiddenItems bool          `json:"showHiddenItems"`
	Name            string        `json:"name"`
	Names           []string      `json:"names"`
	Data            []FileDetails `json:"data"`
	NewName         string        `json:"newName"`
	TargetPath      string        `json:"targetPath"`
	RenameFiles     []string      `json:"renameFiles"`
}

type FileDownloadReqBody struct {
	DownloadInput string `json:"downloadInput"`
}

type FileDetails struct {
	Name          string `json:"name"`
	Size          string `json:"size"`
	IsFile        bool   `json:"isFile"`
	DateModified  string `json:"dateModified"`
	DateCreated   string `json:"dateCreated"`
	Type          string `json:"type"`
	FilterPath    string `json:"filterPath"`
	Permission    string `json:"permission"`
	HasChild      bool   `json:"hasChild"`
	Location      string `json:"location"`
	MultipleFiles bool   `json:"multipleFiles"`
}

type DownloadObj struct {
	Action        string        `json:"action"`
	Path          string        `json:"path"`
	Names         []string      `json:"names"`
	Data          []FileDetails `json:"data"`
	Authorization string        `json:"authorization"`
}

type FileUploadReqBody struct {
	Path     string `json:"path"`
	Action   string `json:"action"`
	Filename string `json:"filename"`
	Data     string `json:"data"`
}

type FileStruct struct {
	Name          string    `json:"name"`
	Size          string    `json:"size"`
	IsFile        bool      `json:"isFile"`
	DateModified  time.Time `json:"dateModified"`
	DateCreated   time.Time `json:"dateCreated"`
	Type          string    `json:"type"`
	FilterPath    string    `json:"filterPath"`
	Permission    any       `json:"permission"`
	HasChild      bool      `json:"hasChild"`
	Location      string    `json:"location"`
	MultipleFiles bool      `json:"multipleFiles"`
}
