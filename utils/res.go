package utils

import (
	"encoding/json"
	"net/http"
	"api-restful-cake-store/models"
)

type ResponseList struct {
	Status string `json:"status"`
	Data []models.Cake `json:"data"`
}

type Response struct {
	Status string `json:"status"`
	Data []models.Cake `json:"data"`
}

type ResponseAction struct {
	Status string `json:"status"`
	Description string `json:"description"`
}


func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	ubahkeByte, err := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "error om", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(ubahkeByte))
}
