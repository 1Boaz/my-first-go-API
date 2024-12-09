package main

import (
	"fmt"
	"net/http"

	"github.com/1Boaz/my-first-go-API/internal/handelers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handelers.Handler(r)

	fmt.Println("Starting GO API service...")

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil{
		log.Error(err)
	}
}