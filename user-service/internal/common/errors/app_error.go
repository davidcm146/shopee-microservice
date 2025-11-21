package Errors

import "fmt"

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func BadRequest(msg string) *AppError {
	return &AppError{
		Code:    "BAD_REQUEST",
		Message: msg,
		Status:  400,
	}
}

func Unauthorized(msg string) *AppError {
	return &AppError{
		Code:    "UNAUTHORIZED",
		Message: msg,
		Status:  401,
	}
}

func NotFound(msg string) *AppError {
	return &AppError{
		Code:    "NOT_FOUND",
		Message: msg,
		Status:  404,
	}
}

func Internal(msg string) *AppError {
	return &AppError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: msg,
		Status:  500,
	}
}

func Forbidden(msg string) *AppError {
	return &AppError{
		Code:    "FORBIDDEN",
		Message: msg,
		Status:  403,
	}
}

func UnprocessableEntity(msg string) *AppError {
	return &AppError{
		Code:    "UNPROCESSABLE_ENTITY",
		Message: msg,
		Status:  422,
	}
}
