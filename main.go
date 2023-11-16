package main

import (
	"golang_mvc_REST_API/controllers"
	"golang_mvc_REST_API/db"
	"golang_mvc_REST_API/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	menu1 := &models.Menu{MenuItems: []models.MenuItem{{Name: "Борщ", Price: 10}}}
	db := db.NewInMemoryState()

	r := mux.NewRouter()

	showMenuHandler := func(w http.ResponseWriter, r *http.Request) {
		controllers.ShowMenuController(w, r, menu1)
	}
	makeOrderHandler := func(w http.ResponseWriter, r *http.Request) {
		controllers.MakeOrderController(w, r, db)
	}
	deleteOrderHandler := func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteOrderController(w, r, db)
	}
	example := func(w http.ResponseWriter, r *http.Request) {
		controllers.Example(w, r)
	}

	r.HandleFunc("/menu", showMenuHandler).Methods("GET")
	r.HandleFunc("/order/make", makeOrderHandler).Methods("POST")
	r.HandleFunc("/order/delete", deleteOrderHandler).Methods("POST")
	r.HandleFunc("/order/example", example).Methods("GET")

	http.ListenAndServe("localhost:8080", r)
}
