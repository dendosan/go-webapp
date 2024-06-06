package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
    maxOpenConns = 25
    maxIdleConns = 25
    connMaxLifetime = 5 * time.Minute
)

func initMySQLDb(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // test db conn
    if err = db.Ping(); err != nil {
        return nil, err
    }

    db.SetMaxOpenConns(maxOpenConns)
    db.SetMaxIdleConns(maxIdleConns)
    db.SetConnMaxLifetime(connMaxLifetime)

    return db, nil
}
