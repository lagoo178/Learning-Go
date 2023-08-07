package main
import (
	"fmt"
   )
func main() {
	// Working with variadic functions
	fmt.Println(addNums(1,2,3,4,5)) // 15
	fmt.Println(addNums(1,2,3)) // 6
}

func addNums(nums ... int) int {
	total := 0
	for _, n := range nums {
	total += n
	}
	return total
}