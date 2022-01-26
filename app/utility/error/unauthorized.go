package error

type ErrUnauthorized struct {
	Message string
}

func (e ErrUnauthorized) Error() string {
	return e.Message
}

func NewErrUnauthorized() ErrUnauthorized {
	return ErrUnauthorized{
		Message: "unauthorized request",
	}
}
