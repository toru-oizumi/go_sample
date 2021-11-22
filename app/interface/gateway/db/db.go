package db

type DBService interface {
	IsDuplicateError(error) bool
}
