package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MachadoMichael/translator/model"
	"github.com/MachadoMichael/translator/service"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody model.TranslationInfo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	log.Printf("Received request: %+v", reqBody)

	translated, err := service.GoogleTranslate(reqBody)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, translated)

}
