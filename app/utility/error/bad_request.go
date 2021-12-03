package error

type ErrBadRequest struct {
	Message string
}

func (e ErrBadRequest) Error() string {
	return e.Message
}

func NewErrBadRequest(message string) ErrBadRequest {
	return ErrBadRequest{
		Message: message,
	}
}
