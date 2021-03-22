package xmysql

import (
	"database/sql"
	"fmt"
	"time"
	"weather_mgr/pkg/xzap"

	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysql(dsn string) (*sql.DB, error) {
	if dsn == "" {
		xzap.Error("dsn nil")
		return nil, errors.New(fmt.Sprintf("db config nil"))
	}

	handler, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	handler.SetMaxOpenConns(10)
	handler.SetMaxIdleConns(10)
	handler.SetConnMaxLifetime(2 * time.Minute)
	err = handler.Ping()
	return handler, nil
}
