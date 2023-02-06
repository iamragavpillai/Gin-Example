package main

import (
	"Gin-example/controller"
	"Gin-example/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	godotenv.Load(".env")

	DB := models.ConnectDB()
	c := controller.NewConnection(DB)
	router := mux.NewRouter()

	router.HandleFunc("/book", c.AddBook).Methods(http.MethodPost)
	http.ListenAndServe(":4000", router)

}
