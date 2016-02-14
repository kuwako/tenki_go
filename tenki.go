package main

import (
	"fmt"
	"github.com/m0a/easyjson"
	"net/http"
)

// 東京の天気を取得するAPI
var tenki_url = "http://weather.livedoor.com/forecast/webservice/json/v1?city=130010"

func main() {
	// urlにアクセス
	resp, err := http.Get(tenki_url)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(resp.StatusCode)
	}

	// jsonパース
	json, err := easyjson.NewEasyJson(resp.Body)
	if err != nil {
		panic("json convert err")
	}

	fmt.Println(json.K("description").K("text"))
	// 条件次第でpepper_botに投げる

}
