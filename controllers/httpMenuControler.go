package controllers

import (
	"golang_mvc_REST_API/models"
	"golang_mvc_REST_API/views"
	"net/http"
)

func ShowMenuController(w http.ResponseWriter, r *http.Request, m *models.Menu) {
	w.Write(views.JsonMenu(m))
}
