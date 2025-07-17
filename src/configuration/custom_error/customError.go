package customError

type CustomError struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int64    `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewCustomErr(message string, err string, code int32, causes []Causes) &CustomError {
	return &CustomError{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func (e *CustomError) Error() string {
	return e.Message
}
func NewBadRequestError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}