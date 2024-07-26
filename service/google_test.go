package service

import (
	"testing"

	"github.com/MachadoMichael/translator/model"
	"github.com/stretchr/testify/assert"
)

func TestGoogleTranslate(t *testing.T) {
	ti := model.TranslationInfo{
		Text: "Hello",
		SL:   "en",
		TL:   "pt",
	}

	translatedText, err := GoogleTranslate(ti)
	assert.NoError(t, err)
	assert.Equal(t, "Olá", translatedText)
}
