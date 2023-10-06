package main
import (
	"fmt"
	"sort"
)

// Creating Maps in Go


func main() {
	
	// heights["Peter"] = 170
	// heights["Joan"] = 168
	// heights["Jan"] = 175

	// Initializing a map with a map literal
	heights := map[string]int{
		"Peter": 170,
		"Joan": 168,
		"Jan": 175, // <-- note the comma here
	}

	// Iterating over a map
	for k, v := range heights {
		fmt.Println(k, v)
	}

	// get all the keys in map
	var keys []string
	for k := range heights {
		keys = append(keys, k)
	}
	fmt.Println(keys) // [Jan Peter Joan]

	// Setting the iteration order in a map
	sort.Strings(keys)
 	fmt.Println(keys) // [Jan Joan Peter]

	// After the keys are sorted, you can then use the for-range loop to iterate over the 
	// keys and print out the value of each key
	for _, k := range keys {
		fmt.Println(k, heights[k])
	}
}