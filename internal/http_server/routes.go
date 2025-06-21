package http_server

import (
	"clam_crawler_av/internal/http_server/ccav"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewServiceRouter() *mux.Router {
	r := mux.NewRouter()

	//Route Register
	r.Handle("/", http.RedirectHandler("/", http.StatusFound))
	r.HandleFunc("/ccav", ccav.HealthCheck).Methods("GET")
	r.HandleFunc("/ccav/{id}", ccav.FileUploadHandler).Methods("GET")

	log.Println("API v1 routes registered")
	return r
}
