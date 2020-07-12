package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	foodUrl, _ = url.Parse("https://api.nal.usda.gov/fdc/v1")
	apiKey string
)

func FoodsSearch(keywords []string) string {
	apiKeyCheck()
	apiKey = os.Getenv("API_KEY")
	foodUrl.Path += "/foods/search"
	params := url.Values{}
	params.Add("api_key", apiKey)
	params.Add("query", buildSearchString(keywords))
	// TODO: make these params configurable
	params.Add("pageSize", "1")
	params.Add("dataType", "Foundation")
	params.Add("requireAllWords", "true")
	foodUrl.RawQuery = params.Encode()
	resp, respErr := http.Get(foodUrl.String())
	if respErr != nil {
		log.Fatal(respErr)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Query returned a HTTP ", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	// TODO: figure out unmarshalling json
	return string(body)
	//var s string
	//jsonErr := json.Unmarshal(body, &s)
	//if jsonErr != nil {
	//	log.Fatal(jsonErr)
	//}
	//return s
}

func apiKeyCheck() {
	_, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatal("API_KEY not set! Cancelling search...\n")
	}
}

func buildSearchString(keywords []string) string {
	var builder strings.Builder
	if i := len(keywords); i == 1 {
		builder.WriteString(keywords[0])
	} else {
		for i := 0; i < len(keywords); i++ {
			if i != (len(keywords) - 1) {
				stringWithSpace := fmt.Sprintf("%s ")
				builder.WriteString(stringWithSpace)
			} else {
				builder.WriteString(keywords[i])
			}
		}
	}
	return builder.String()
}
