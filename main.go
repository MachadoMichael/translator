package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MachadoMichael/translator/translator"
)

// Define the requestBody struct
type requestBody struct {
	Text string `json:"text"`
	SL   string `json:"sl"`
	TL   string `json:"tl"`
}

// Handler function to handle the POST request
func translateHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body
	var reqBody requestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Log the received data (for demonstration purposes)
	log.Printf("Received request: %+v", reqBody)

	// Respond with a success message
	translated, err := translator.Translate(reqBody.Text, reqBody.SL, reqBody.TL)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, translated)

}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register the handler function for the /translate path
	mux.HandleFunc("/translate", translateHandler)

	// Start the server
	log.Println("Server is starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
