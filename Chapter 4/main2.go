package main

// import "fmt"

func main() {
	//fibonacci sequence
	max := 100
	a, b := 0, 1
	for b <= max {
		println(b)
		a, b = b, a+b
	}
}