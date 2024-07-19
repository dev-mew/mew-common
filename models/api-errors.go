package models

type ApiError struct {
	Error ErrorContent `json:"error"`
}

type ErrorContent struct {
	Status  int    `json:"status"`
	Name    string `json:"name"`
	Message string `json:"message"`
}
