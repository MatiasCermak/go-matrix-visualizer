package main

import (
	"log"
	"matrix-visualizer/backend/pkg/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/matrix/{matrixOperation}", handlers.GenerateMatrix).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
