package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL ドライバのインポート
)

func InitDB(dataSourceName string) (*sql.DB, error) {
	// DBの接続をオープン
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// PostgreSQLに接続できるか確認
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	// 成功したらDBインスタンスを返す
	return db, nil
}
