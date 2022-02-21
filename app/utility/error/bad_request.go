package error

type ErrBadRequest struct {
	customError
}

func NewErrBadRequest(message string) ErrBadRequest {
	return ErrBadRequest{
		customError{
			Title:   "bad_request",
			Message: message,
		},
	}
}
