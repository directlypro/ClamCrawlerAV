package ccav

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "File Upload()\n")

	//1. Parameter input for multipart form file
	r.ParseMultipartForm(200 << 20)

	//2. Retrieve file from form-data.
	file, handler, err := r.FormFile("file")
	fileName := r.FormValue("filename")
	if err != nil {
		errStr := fmt.Sprintf("Error uploading file: %v", err)
		fmt.Printf(errStr)
		fmt.Fprintf(w, errStr)
		return
	}
	defer file.Close()

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

}
