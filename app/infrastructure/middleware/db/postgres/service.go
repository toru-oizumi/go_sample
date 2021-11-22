package postgres

import (
	"go_sample/app/interface/gateway/db"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type PostgresSqlDBService struct{}

func (s *PostgresSqlDBService) IsDuplicateError(err error) bool {
	pqErr := err.(*pgconn.PgError)
	return pqErr.Code == pgerrcode.UniqueViolation
}

func NewDBService() db.DBService {
	return &PostgresSqlDBService{}
}
