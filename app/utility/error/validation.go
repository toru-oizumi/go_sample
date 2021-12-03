package error

type ErrValidationError struct {
	Message string
}

func (e ErrValidationError) Error() string {
	return e.Message
}

func NewErrValidationError(err error) ErrValidationError {
	return ErrValidationError{
		Message: err.Error(),
	}
}
