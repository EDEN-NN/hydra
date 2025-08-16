package apperrors

import "fmt"

type ErrorCode string

const (
	ECONFLICT     ErrorCode = "CONFLICT"
	ENOTFOUND     ErrorCode = "NOT_FOUND"
	EINVALID      ErrorCode = "INVALID"
	EUNAUTHORIZED ErrorCode = "UNAUTHORIZED"
	EINTERNAL     ErrorCode = "INTERNAL"
)

type AppError struct {
	Code     ErrorCode              `json:"code"`
	Message  string                 `json:"message"`
	Err      error                  `json:"errorDetail"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v (code: %s)", e.Message, e.Err, e.Code)
	}
	return fmt.Sprintf("")
}

func NewError(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewConflictError(entity string, err error) *AppError {
	return NewError(ECONFLICT, fmt.Sprintf("%s field error", entity), err)
}

func NewNotFoundError(entity string) *AppError {
	return NewError(ENOTFOUND, fmt.Sprintf("%s entity or field not found", entity), nil)
}
