package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenDB(sourceName, user, password, database string) {
	dbDataSource := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", user, password, sourceName, database)
	initDB, err := sql.Open("mysql", dbDataSource)
	errorkit.ErrorHandled(err)

	DB = initDB
}
