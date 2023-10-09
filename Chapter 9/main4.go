package main
import (
	"encoding/json"
	"fmt"
)
// Mapping custom attribute names
type Rates struct {
	Base string `json:"base currency"`
	Symbol string `json:"destination currency"`
}
func main() {
	jsonString :=
					`{
						"base currency":"EUR",
						"destination currency":"USD"
					}`
	var rates Rates
	json.Unmarshal([]byte(jsonString), &rates)
	fmt.Println(rates.Base) // EUR
	fmt.Println(rates.Symbol) // USD
}