package dbconn

import (
	"database/sql"
	"fmt"

	"github.com/Karanth1r3/ws_chat_learn/config"
)

func ConnectDB(cfg config.DB) (*sql.DB, error) {

	dbConnStr := fmt.Sprintf(
		"host=%s port =%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.Username,
		cfg.Password,
	)
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		return nil, fmt.Errorf("connection to db failed: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(0)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db is not available %w", err)
	}

	return db, nil
}
