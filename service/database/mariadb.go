package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db *sql.DB
}

func NewStore(username string, password string, hostname string, port int, database string) (*Store, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, database)
	db, err := sql.Open("mysql", connection)

	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
