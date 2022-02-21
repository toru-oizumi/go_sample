package error

type customError struct {
	Title   string
	Message string
}

func (e customError) Error() string {
	return e.Message
}
