package error

type ErrEntityAlreadyExists struct {
	customError
}

func NewErrEntityAlreadyExists() ErrEntityAlreadyExists {
	return ErrEntityAlreadyExists{
		customError{
			Title:   "entity_already_exists",
			Message: "this entity already exists",
		},
	}
}

type ErrEmailAlreadyExists struct {
	customError
}

func NewErrEmailAlreadyExists(message string) ErrEmailAlreadyExists {
	return ErrEmailAlreadyExists{
		customError{
			Title:   "email_already_exists",
			Message: message,
		},
	}
}
