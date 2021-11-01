package error

type ErrUnknown struct {
	Message string
}

func (e *ErrUnknown) Error() string {
	return e.Message
}

func NewErrUnknown() *ErrUnknown {
	return &ErrUnknown{
		Message: "unknown error",
	}
}
