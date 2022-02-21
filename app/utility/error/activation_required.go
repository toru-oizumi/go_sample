package error

type ErrActivationRequired struct {
	customError
}

func NewErrActivationRequired() ErrActivationRequired {
	return ErrActivationRequired{
		customError{
			Title:   "activation_required",
			Message: "activation required",
		},
	}
}
