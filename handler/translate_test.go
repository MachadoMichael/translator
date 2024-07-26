package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MachadoMichael/translator/model"
	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	reqBody := model.TranslationInfo{
		Text: "Hello",
		SL:   "en",
		TL:   "pt",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/translate", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Translate)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Olá")
}

func TestTranslateMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("GET", "/translate", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Translate)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	assert.Contains(t, rr.Body.String(), "Method not allowed")
}
