package main

import (
	"fmt"
	"net/http"

	"goapi/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	r := chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API services")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
