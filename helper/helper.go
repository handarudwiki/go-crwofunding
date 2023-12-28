package helper

import "github.com/go-playground/validator/v10"

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}) response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatError(err error) []string {
	var errors []string

	validationErr, ok := err.(validator.ValidationErrors)
	if !ok {
		// Handle jenis error yang tidak diharapkan di sini
		errors = append(errors, "Error: Unexpected validation error type")
		return errors
	}

	for _, e := range validationErr {
		errors = append(errors, e.Error())
	}
	return errors
}
