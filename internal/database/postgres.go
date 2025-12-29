package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ineoo/go-planigramme/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	maxConn, _ := strconv.Atoi(config.Config("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(config.Config("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(config.Config("DB_MAX_LIFETIME_CONNECTIONS"))

	db, err := sqlx.Connect("pgx", config.Config("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(maxLifetimeConn))

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}
