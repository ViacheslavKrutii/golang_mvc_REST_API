package db

import (
	"golang_mvc_REST_API/models"
	"net/http"
	"slices"
)

type InMemoryState struct {
	orders map[models.User][]models.Order
}

func NewInMemoryState() *InMemoryState {
	return &InMemoryState{orders: make(map[models.User][]models.Order)}
}

func (i InMemoryState) AddOrder(w http.ResponseWriter, r *http.Request, newOrder models.Order) {
	i.orders[newOrder.User] = append(i.orders[newOrder.User], newOrder)
}

func (i InMemoryState) DeleteOrder(w http.ResponseWriter, r *http.Request, newDeleteRequest models.DeleteOrderRequest) {
	ordersSlice, ok := i.orders[newDeleteRequest.User]
	if !ok {
		http.Error(w, "You have not order.", http.StatusBadRequest)
		return
	}
	i.orders[newDeleteRequest.User] = slices.Delete(ordersSlice, newDeleteRequest.IdOrder, newDeleteRequest.IdOrder)
}