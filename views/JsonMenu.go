package views

import (
	"Proj/golang_mvc_REST_API/models"
	"encoding/json"
)

func JsonMenu(m *models.Menu) []byte {
	serialized, _ := json.Marshal(m)
	return serialized
}
