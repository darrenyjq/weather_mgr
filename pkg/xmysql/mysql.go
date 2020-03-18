package xmysql

import (
	"database/sql"
	"fmt"
	"time"

	"errors"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

var dbHandle *sql.DB

func NewMysql(key string) (*sql.DB, error) {
	if dbHandle != nil {
		return dbHandle,nil
	}
	if !viper.IsSet(key) {
		return nil,errors.New(fmt.Sprintf("db config nil: %", key))
	}
	dbHandle, err := sql.Open("mysql", viper.GetString(key+".dsn"));
	if err != nil {
		return nil,err
	}
	dbHandle.SetMaxOpenConns(10)
	dbHandle.SetMaxIdleConns(10)
	dbHandle.SetConnMaxLifetime(1 * time.Minute)
	err = dbHandle.Ping()
	return dbHandle,nil
}

