package error

type ErrEntityNotExists struct {
	customError
}

func NewErrEntityNotExists(entityName string) ErrEntityNotExists {
	return ErrEntityNotExists{
		customError{
			Title:   "entity_not_exists",
			Message: "this '" + entityName + "' does not exist",
		},
	}
}
