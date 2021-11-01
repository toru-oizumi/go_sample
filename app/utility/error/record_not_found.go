package error

type ErrRecordNotFound struct {
	Message string
}

func (e *ErrRecordNotFound) Error() string {
	return e.Message
}

func NewErrRecordNotFound() *ErrRecordNotFound {
	return &ErrRecordNotFound{
		Message: "this record not found",
	}
}
