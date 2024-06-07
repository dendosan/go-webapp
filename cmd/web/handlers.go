package main

import (
	"fmt"
	"go-webapp/pets"
	"net/http"
	"slices"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

var pages = []string {
    "test",
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

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
    var t toolbox.Tools
    _ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}
func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
    var t toolbox.Tools
    _ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) TestPatterns (w http.ResponseWriter, r *http.Request) {
    app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory (w http.ResponseWriter, r *http.Request) {
    var t toolbox.Tools
    dog, err := pets.NewPetFromAbstractFactory("dog")
    if err != nil {
        _ = t.ErrorJSON(w, err, http.StatusBadRequest)
        return
    }
    _ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory (w http.ResponseWriter, r *http.Request) {
    var t toolbox.Tools
    cat, err := pets.NewPetFromAbstractFactory("cat")
    if err != nil {
        _ = t.ErrorJSON(w, err, http.StatusBadRequest)
        return
    }
    _ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
    dogBreeds, err := app.Models.DogBreed.All()
    var t toolbox.Tools

    if err != nil {
        _ = t.ErrorJSON(w, err, http.StatusBadRequest)
        return
    }
    _ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}
