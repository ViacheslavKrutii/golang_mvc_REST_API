package controllers

import (
	"encoding/json"
	"golang_mvc_REST_API/db"
	"golang_mvc_REST_API/models"
	"golang_mvc_REST_API/views"
	"io"
	"log"
	"net/http"
)

func MakeOrderController(w http.ResponseWriter, r *http.Request, db db.DbInterface) {
	log.Println("Make order controller")
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	var newOrder models.Order

	err = json.Unmarshal(body, &newOrder)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	db.AddOrder(w, r, newOrder)

}

func DeleteOrderController(w http.ResponseWriter, r *http.Request, db db.DbInterface) {
	log.Println("Delete order controller")
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	var newDeleteRequest models.DeleteOrderRequest

	err = json.Unmarshal(body, &newDeleteRequest)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	db.DeleteOrder(w, r, newDeleteRequest)
}

func Example(w http.ResponseWriter, r *http.Request) {

	orderExample := models.Order{OrderBody: map[string]int{"Борщ": 3}, User: models.User{Name: "me"}}

	w.Write(views.JsonOrderExample(&orderExample))
}
