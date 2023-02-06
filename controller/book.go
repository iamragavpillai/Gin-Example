package controller

import (
	"Gin-example/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handlerCustom) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		panic(err)
	}

	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
