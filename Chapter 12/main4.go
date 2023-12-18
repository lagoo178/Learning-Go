package main

import (
	"fmt"
	"time"
)

// Asynchronously Waiting on Channels

func fib(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
		time.Sleep(2 * time.Second)
	}
	close(c)
}

func counter(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	close(c)
}
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go fib(10, c1) // generate 10 fibo nums
	go counter(10, c2) // generate 10 numbers
	for i := range c1 {
		fmt.Println("fib()", i)
	}
	for i := range c2 {
		fmt.Println("counter()", i)
	}
}