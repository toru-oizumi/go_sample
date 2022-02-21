package error

type ErrAuthenticationFailed struct {
	customError
}

func NewErrAuthenticationFailed() ErrAuthenticationFailed {
	return ErrAuthenticationFailed{
		customError{
			Title:   "authentication_failed",
			Message: "authentication_failed",
		},
	}
}
