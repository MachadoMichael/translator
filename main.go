package main

import (
	"log"
	"net/http"

	"github.com/MachadoMichael/translator/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/translate", handler.Translate)

	log.Println("Server is starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
