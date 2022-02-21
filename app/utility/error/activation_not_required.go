package error

type ErrActivationNotRequired struct {
	customError
}

func NewErrActivationNotRequired() ErrActivationNotRequired {
	return ErrActivationNotRequired{
		customError{
			Title:   "activation_not_required",
			Message: "activation not required",
		},
	}
}
