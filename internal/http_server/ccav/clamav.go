package ccav

import (
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	log.Println("Health Check")
	_, err := w.Write([]byte(`{"alive": true}`))
	if err != nil {
		return
	}
}
