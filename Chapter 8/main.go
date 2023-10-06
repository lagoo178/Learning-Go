package main
import (
	"fmt"
	
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

	// Checking the existence of a key
	if v, ok := heights["Jim"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist")
	}
	fmt.Println(heights["Peter"]) // 170
	fmt.Println(heights["Jan"]) // 175

	// Deleting a key
	if _, ok := heights["Joan"]; ok {
		delete(heights, "Joan")
		fmt.Println("Key successfully deleted")
	} else {
		fmt.Println("Key does not exist")
	}
	// fmt.Println(heights["Joan"]) // 168

	// Getting the number of items in a map
	var weights map[string]int
	fmt.Println(len(weights)) // 0
}