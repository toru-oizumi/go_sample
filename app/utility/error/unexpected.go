package error

type ErrUnexpected struct {
	customError
}

func NewErrUnexpected(message string) ErrUnexpected {
	return ErrUnexpected{
		customError{
			Title:   "unexpected",
			Message: message,
		},
	}
}
