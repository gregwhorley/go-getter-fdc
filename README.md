# go-getter-fdc
Simple CLI tool that retrieves food data from USDA's FoodData Central API

https://fdc.nal.usda.gov/api-guide.html

You will need an API key from FDC to use this app: https://fdc.nal.usda.gov/api-key-signup.html

## Building
`go build`

## Testing
`go test -v`

## Usage
Search for food by keywords:
`API_KEY=yourKey ./go-getter-fdc search cheddar cheese`
