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

// URL query params defined in FoodsSearch() for now
type FoodSearchCriteria struct {
	DataType           []string
	Query              string
	GeneralSearchInput string
	PageNumber         int
	RequireAllWords    bool
}

type FoodNutrients struct {
	NutrientId            int
	NutrientNumber        string
	NutrientName          string
	UnitName              string
	DerivationCode        string
	DerivationDescription string
	Value                 float32
}

type Foods struct {
	FdcId                  int
	DataType               string
	Description            string
	FoodCode               string
	FoodNutrients          []FoodNutrients
	PublishedDate          string
	ScientificName         string
	BrandOwner             string
	GtinUpc                string
	Ingredients            string
	NdbNumber              string
	AdditionalDescriptions string
	AllHighlightFields     string
	Score                  float32
}

type FoodsSearchJson struct {
	FoodSearchCriteria FoodSearchCriteria
	TotalHits          int
	CurrentPage        int
	TotalPages         int
	Foods              []Foods
}
