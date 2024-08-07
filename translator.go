package translator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TranslationInfo struct {
	Text string `json:"text"`
	SL   string `json:"sl"`
	TL   string `json:"tl"`
}

func Translate(ti TranslationInfo) (string, error) {
	txt := strings.Replace(ti.Text, " ", "%20", -1)
	url := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" + ti.SL + "&tl=" + ti.TL + "&dt=t&q=" + txt

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	var result []interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	str := strings.Split(string(body), "\"")
	return str[1], nil
}
