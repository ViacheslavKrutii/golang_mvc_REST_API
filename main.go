package main

import (
	"golang_mvc_REST_API/controllers"
	"golang_mvc_REST_API/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	i1 := models.MenuItem{Name: "Борщ", Price: 10}
	m := &models.Menu{MenuItems: []models.MenuItem{i1}}

	r := mux.NewRouter()

	showMenuHandler := func(w http.ResponseWriter, r *http.Request) {
		controllers.ShowMenuController(w, r, m)
	}

	r.HandleFunc("/menu", showMenuHandler).Methods("GET")
	r.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")

	http.ListenAndServe("localhost:8080", r)
}
