package mysql

import (
	"go_sample/app/interface/gateway/db"

	"github.com/go-sql-driver/mysql"
)

type MySqlDBService struct{}

func (s *MySqlDBService) IsDuplicateError(err error) bool {
	mysqlErr := err.(*mysql.MySQLError)
	// https://dev.mysql.com/doc/refman/5.6/ja/error-messages-server.html
	// エラー: 1062 SQLSTATE: 23000 (ER_DUP_ENTRY)
	// メッセージ: '%s' はキー %d で重複しています
	return mysqlErr.Number == 1062
}

func NewDBService() db.DBService {
	return &MySqlDBService{}
}
