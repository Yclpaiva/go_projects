package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"1stgame_backend/internal"
)

type ApiResponse struct {
	Message string `json:"message"`
}

func resourcesDataHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header to allow all origins
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Set the response header to be JSON
	w.Header().Set("Content-Type", "application/json")

	resourcesMap := internal.DbQueryResources()
	json.NewEncoder(w).Encode(resourcesMap)
}

func industriesConsumptionDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	industriesConsumption := internal.DbQueryIndustriesConsumption()
	json.NewEncoder(w).Encode(industriesConsumption)
}

func industriesProductionDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	industriesProduction := internal.DbQueryIndustriesProduction()
	json.NewEncoder(w).Encode(industriesProduction)
}

func main() {
	println("Server starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/api/resources", resourcesDataHandler)
	http.HandleFunc("/api/industries/consumption", industriesConsumptionDataHandler)
	http.HandleFunc("/api/industries/production", industriesProductionDataHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		println("Server failed to start!")
	}
}
