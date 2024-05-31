package main

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/go-chi/chi/v5"
)

var pages = []string {
    "home",
    "about",
    "cat-breed",
    "cat-breeds",
    "cat-breeders",
    "dog-breed",
    "dog-breeds",
    "dog-breeders",
}

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
    app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
    page := chi.URLParam(r, "page")
    if !slices.Contains(pages, page) {
        fmt.Printf("Requested page: %s is not a valid page.\n", page)
        return
    }
    app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}
