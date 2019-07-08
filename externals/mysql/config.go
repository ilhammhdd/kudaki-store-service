package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	_ "github.com/go-sql-driver/mysql"
)

var CommandDB *sql.DB
var QueryDB *sql.DB

func OpenDB(sourceName, user, password, database string) *sql.DB{
	dbDataSource := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", user, password, sourceName, database)
	initDB, err := sql.Open("mysql", dbDataSource)
	errorkit.ErrorHandled(err)

	return initDB
}
