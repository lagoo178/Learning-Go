package main
import (
	"fmt"
   )
func main() {
	// Returning values from functions
	o, e := countOddEven("12345")
	fmt.Println(o,e) // 3 2
}

func countOddEven(s string) (int,int) {
	odds, evens := 0,0
	for _, c := range s {
	if int(c) % 2 == 0 {
	evens++
	} else {
	odds++
	}
	}
	return odds,evens
}