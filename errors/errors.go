package errors

type ServerError struct {
	Message string `json:"message"`
}

type NotFoundError struct {
	Message string `json:"message"`
}

type BadRequestError struct {
	Message string `json:"message"`
}
