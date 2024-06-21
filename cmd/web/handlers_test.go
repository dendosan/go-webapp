package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsJSON(t *testing.T) {
	// create request
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)

	// create response recorder
	rr := httptest.NewRecorder()

	// create handler
	handler := http.HandlerFunc(testApp.GetAllDogBreedsJSON)

	// serve handler
	handler.ServeHTTP(rr, req)

	// check response
	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code; got %d, wanted 200", rr.Code)
	}
}

func TestApplication_GetAllCatBreeds(t *testing.T) {
	// create request
	req, _ := http.NewRequest("GET", "/api/cat-breeds", nil)

	// create response recorder
	rr := httptest.NewRecorder()

	// create handler
	handler := http.HandlerFunc(testApp.GetCatBreeds)

	// serve handler
	handler.ServeHTTP(rr, req)

	// check response
	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code; got %d, wanted 200", rr.Code)
	}
}
