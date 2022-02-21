package error

type ErrChangePasswordRequired struct {
	customError
}

func NewErrChangePasswordRequired() ErrChangePasswordRequired {
	return ErrChangePasswordRequired{
		customError{
			Title:   "change_password_required",
			Message: "change password required",
		},
	}
}
