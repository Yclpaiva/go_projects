package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"1stgame_backend/internal"
)

type ApiResponse struct {
	Message string `json:"message"`
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header to allow all origins
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Set the response header to be JSON
	w.Header().Set("Content-Type", "application/json")
	// returned := internal.DbQuerytest()
	// response := ApiResponse{Message: strings.Join(returned, ", ")

	resourcesMap := internal.DbQueryResources()
	var returned []string
	for key, value := range resourcesMap {
		returned = append(returned, fmt.Sprintf("%s: %.2f", key, value))
	}
	response := ApiResponse{Message: strings.Join(returned, ", ")}
	json.NewEncoder(w).Encode(response)
}

func main() {
	println("Server starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/api/", dataHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		println("Server failed to start!")
	}
}
