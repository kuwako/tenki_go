package main

import (
	"fmt"
	"github.com/m0a/easyjson"
	"net/http"
)

// 東京の天気を取得するAPI
var tenki_url = "http://weather.livedoor.com/forecast/webservice/json/v1?city=130010"
var tenki_text = ""

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

	fmt.Println(json.K("title"))

	for _, content := range json.K("forecasts").RangeObjects() {
		dateLabel, _ := content.K("dateLabel").AsString()
		if dateLabel == "今日" {
			today, _ := content.K("date").AsString()
			weather, _ := content.K("telop").AsString()
			description, _ := json.K("description").K("text").AsString()

			tenki_text += today + "の東京の天気は " + weather + " です\n\n" + description
		}
	}

	fmt.Println(tenki_text)

	// 条件次第でpepper_botに投げる
}
