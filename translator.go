package translator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Translater(txt string, sl string, tl string) {
	txt = strings.Replace(txt, " ", "%20", -1)
	url := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" + sl + "&tl=" + tl + "&dt=t&q=" + txt

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
		return
	}

	str := strings.Split(string(body), "\"")
	fmt.Println(str[1])
}
