package main
import (
	"fmt"
   )
func main() {
	// Passing arguments by value and by pointer
	x := 5
	y := 6
	swap(x, y)
	fmt.Println(x, y) // 5 6
}

func swap(a, b int) {
	a, b = b, a
}