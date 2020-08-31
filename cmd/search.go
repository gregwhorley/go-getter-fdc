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
package cmd

import (
	"fmt"
	"github.com/gregwhorley/go-getter-fdc/pkg/client"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var (
	searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for foods by keywords",
	Long: `This action and subsequent arguments are passed in to the FoodData Central API 

Example: ./go-getter-fdc search onion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Search called for %v...\n", args)
		queryOptions = client.QueryOptionsFiller(pageSize, dataType, requireAllWords)
		foodsSearch := client.FoodsSearch(args, queryOptions)
		for _, food := range foodsSearch.Foods {
			fmt.Printf("Basic Data:\n")
			fmt.Printf("  Description: %v\n", food.Description)
			fmt.Printf("  Data Type: %v\n", food.DataType)
			fmt.Printf("  Ingredients: %v\n", food.Ingredients)
			fmt.Printf("Nutrient Data:\n")
			for _, nutrients := range food.FoodNutrients {
				fmt.Printf("  Name: %v\n", nutrients.NutrientName)
				fmt.Printf("  Amount: %v%v\n", nutrients.NutrientNumber, nutrients.UnitName)
			}
		}
	},
}
	queryOptions map[string]string
	pageSize string
	dataType string
	requireAllWords string
)

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	searchCmd.Flags().StringVar(&pageSize,"page-size", "1", "Set the page size for the result set. Defaults to 1.")
	searchCmd.Flags().StringVar(&dataType, "datatype", "Foundation", "Set the datatype to one of the following:\nFoundation\nBranded\nSurvey\nLegacy")
	searchCmd.Flags().StringVar(&requireAllWords, "require-all-words", "true", "Require all keywords in search results. Defaults to true.")
}
