package main
import (
	"encoding/json"
	"fmt"
)

// Decoding JSON to a struct

type People struct {
	Firstname string
	Lastname string
}
func main() {
	var person People
	jsonString := `{
						"firstname":"Wei-Meng",
						"lastname":"Lee"
					}` 
	err := json.Unmarshal([]byte(jsonString), &person)
	if err == nil {
		fmt.Println(person.Firstname)
		fmt.Println(person.Lastname)
	} else {
		fmt.Println(err)
	}
}