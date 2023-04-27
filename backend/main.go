package main

import (
	"log"
	"matrix-visualizer/backend/pkg/handlers"
	"matrix-visualizer/backend/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	middleware.EnableCORS(r)

	r.HandleFunc("/matrix/{matrixOperation}", handlers.GenerateMatrix).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
