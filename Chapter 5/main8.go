package main
import (
	"fmt"
   )

func main() {
	// Implementing closure using anonymous functions
	gen := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
		}
}

func fib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
	f1, f2 = f2, (f1 + f2)
	return f1
	}
}

