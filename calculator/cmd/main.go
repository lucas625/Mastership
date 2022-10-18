package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/lucas625/Mastership/calculator/service/receiver"
)

func createServer() *http.Server {
	service := receiver.New()

	router := mux.NewRouter()
	router.HandleFunc("/add", service.Add)
	router.HandleFunc("/divide", service.Divide)
	router.HandleFunc("/multiply", service.Multiply)
	router.HandleFunc("/subtract", service.Subtract)

	server := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 1 * time.Minute,
		ReadTimeout:  1 * time.Minute,
	}

	fmt.Println("Server running!")
	return server
}

func main() {
	server := createServer()
	err := server.ListenAndServe()
	log.Fatal(err)
}
