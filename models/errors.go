package models

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type ServerError struct {
	Message string `json:"message"`
}
