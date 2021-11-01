package error

import "fmt"

type ErrWsConnectionAlreadyExist struct {
	Message string
}

func (e *ErrWsConnectionAlreadyExist) Error() string {
	return e.Message
}

func NewErrWsConnectionAlreadyExist(objective string, id string) *ErrWsConnectionAlreadyExist {
	return &ErrWsConnectionAlreadyExist{
		Message: fmt.Sprintf("this connection(objective=%s, id=%s) already exists.", objective, id),
	}
}
