package main
import (
	"fmt"
   )

func main() {
	// Declaring an anonymous function
	var i func() int
	i = func() int {
	return 5
	}
	fmt.Println(i()) // 5
}

