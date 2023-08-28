package main
import (
	"fmt"
   )

func main() {
	// Implementing the filter() function using closure
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filter(a,
		func(val int) bool {
			return val%3 == 0
		}
	)
	fmt.Println(evens)
}

func filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for _, v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}
	return result
}

