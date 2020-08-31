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
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	foodUrl, _  = url.Parse("https://api.nal.usda.gov/fdc/v1")
	apiKey      string
	foodResults FoodsSearchJson
)

func FoodsSearch(keywords []string, queryOptions map[string]string) FoodsSearchJson {
	apiKey = GetApiKey()
	foodUrl.Path += "/foods/search"
	params := url.Values{}
	params.Add("api_key", apiKey)
	params.Add("query", BuildSearchString(keywords))
	params.Add("pageSize", queryOptions["pageSize"])
	params.Add("dataType", queryOptions["dataType"])
	params.Add("requireAllWords", queryOptions["requireAllWords"])
	foodUrl.RawQuery = params.Encode()
	resp, respErr := http.Get(foodUrl.String())
	if respErr != nil {
		log.Fatal(respErr)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Expected HTTP 200 but received ", resp.StatusCode)
	}
	defer resp.Body.Close()
	// TODO: I want to omit nutrient data with names like "int:int"
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	if jsonErr := json.Unmarshal(body, &foodResults); jsonErr != nil {
		panic(jsonErr)
	}
	return foodResults
}
