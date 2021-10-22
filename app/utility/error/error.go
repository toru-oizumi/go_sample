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

type ErrRecordDuplicate struct {
	Message string
}

func (e *ErrRecordDuplicate) Error() string {
	return e.Message
}

func NewErrRecordDuplicate() *ErrRecordDuplicate {
	return &ErrRecordDuplicate{
		Message: "this record is duplicated",
	}
}

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
