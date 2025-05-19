package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewForbbidenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewConflictError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Conflict",
		Code:    http.StatusConflict,
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}
