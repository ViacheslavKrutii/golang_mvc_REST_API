package db

import (
	"golang_mvc_REST_API/models"
	"net/http"
)

type DbInterface interface {
	AddOrder(http.ResponseWriter, *http.Request, models.Order)
	DeleteOrder(http.ResponseWriter, *http.Request, models.DeleteOrderRequest)
}
