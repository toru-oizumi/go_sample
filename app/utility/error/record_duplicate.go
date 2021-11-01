package error

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
