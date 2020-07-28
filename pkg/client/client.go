/*
Copyright © 2020 Greg Whorley <greg@whorley.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	foodUrl, _  = url.Parse("https://api.nal.usda.gov/fdc/v1")
	apiKey      string
	foodResults FoodsSearchJson
)

func FoodsSearch(keywords []string) FoodsSearchJson {
	apiKeyCheck()
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
		log.Fatal("Expected HTTP 200 but received ", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	if jsonErr := json.Unmarshal(body, &foodResults); jsonErr != nil {
		panic(jsonErr)
	}
	return foodResults
}

func apiKeyCheck() {
	if _, exists := os.LookupEnv("API_KEY"); !exists {
		log.Fatal("API_KEY not set! Cancelling search...\n")
	} else {
		apiKey = os.Getenv("API_KEY")
	}
}

func buildSearchString(keywords []string) string {
	var builder strings.Builder
	if i := len(keywords); i == 1 {
		builder.WriteString(keywords[0])
	} else {
		for i := 0; i < len(keywords); i++ {
			if i != (len(keywords) - 1) {
				builder.WriteString(fmt.Sprintf("%s ", keywords[i]))
			} else {
				builder.WriteString(keywords[i])
			}
		}
	}
	return builder.String()
}
