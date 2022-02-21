package error

type ErrUnknown struct {
	customError
}

func NewErrUnknown() ErrUnknown {
	return ErrUnknown{
		customError{
			Title:   "unknown",
			Message: "unknown error",
		},
	}
}
