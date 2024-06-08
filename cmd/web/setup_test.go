package main

import (
	"go-webapp/configuration"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
	    App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
