package usecases

import "database/sql"

type DBOperation interface {
	Command(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) (*sql.Row, error)
	QueryRowsToMap(string, ...interface{}) (*[]map[string]interface{}, error)
}
