package db

import (
	"database/sql"
    _ "github.com/lib/pq"
    "fmt"
)

func InitDB(dataSourceName string) (*spl.DB error) {
	db,err := sql.Open("postgres", dataSourceName)
	if err != nill {
		return nil, err
	}
	return db, nil


    // PostgreSQLに接続できるか確認
  err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v". err)
	}

	return db.nil
}