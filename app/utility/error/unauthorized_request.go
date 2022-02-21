package error

type ErrUnauthorizedRequest struct {
	customError
}

func NewErrUnauthorizedRequest() ErrUnauthorizedRequest {
	return ErrUnauthorizedRequest{
		customError{
			Title:   "unauthorized_request",
			Message: "unauthorized request",
		},
	}
}
