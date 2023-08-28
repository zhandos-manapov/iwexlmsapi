package models

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type RespMsg struct {
	Message string `json:"message"`
}
