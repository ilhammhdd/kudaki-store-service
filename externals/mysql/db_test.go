package mysql_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilhammhdd/go-toolkit/errorkit"
)

func TestReplicaDB(t *testing.T) {
	dbDataSource := fmt.Sprintf("%s:%s@%s/?parseTime=true", "root", "mysqlreplicarocks", "tcp(178.62.107.160:3307)")
	db, err := sql.Open("mysql", dbDataSource)
	if errorkit.ErrorHandled(err) {
		t.Error(err)
	}

	err = db.Ping()
	if errorkit.ErrorHandled(err) {
		t.Error(err)
	}
}
