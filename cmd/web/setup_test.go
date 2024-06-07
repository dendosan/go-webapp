package main

import (
	"go-webapp/models"
	"log"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	dsn := "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s"
	db, err := initMySQLDb(dsn)
	if err != nil {
		log.Panic(err)
	}

	testApp = application{
		DBPool: db,
		Models: *models.New(db),
	}

	os.Exit(m.Run())
}