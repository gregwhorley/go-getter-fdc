package client

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
)

var (
	baseUrl = "https://api.nal.usda.gov/fdc/v1"
	apiKey string
)

func getFoodsList(keyword string) string {
	resp, respErr := http.Get(baseUrl+"/foods/search?api_key="+apiKey+"&query="+keyword)
	if respErr != nil || resp.StatusCode != 200 {
		log.Fatal(respErr)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var s string
	jsonErr := json.Unmarshal(body, &s)
	if jsonErr != nil {
		log.Fatal(respErr)
	}
	return s
}
