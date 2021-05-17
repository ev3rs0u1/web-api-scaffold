package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"web-api-scaffold/internal/pkg/config"
	"web-api-scaffold/internal/pkg/logger"
)

func NewConnection() (conn *gorm.DB, err error) {
	var dsn string
	if dsn, err = config.Instance().Database.DSN(); err != nil {
		return
	}

	if conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.NewGORMLogger(),
	}); err != nil {
		return
	}

	var pool *sql.DB
	if pool, err = conn.DB(); err != nil {
		return
	}

	pool.SetMaxIdleConns(15)
	pool.SetMaxOpenConns(30)
	pool.SetConnMaxLifetime(1 * time.Minute)

	return
}
