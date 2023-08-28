package main
import (
 "fmt"
//  "strconv"
)
func main() {
 	// Creating an empty slice
	// s := make([]int, 5)
	// fmt.Println(s)

	r := make([]int, 2, 5)
	fmt.Println(r)
	fmt.Println(len(r)) // 2
	fmt.Println(cap(r)) // 5
}