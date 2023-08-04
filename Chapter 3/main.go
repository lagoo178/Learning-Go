package main

import (
	"fmt"
	// "strconv"
   )

func main() {
	// simple odd and even check
	var num int
	fmt.Print("Please enter your one number: ")
	fmt.Scanf("%d", &num)
	if num % 2 == 1 {
			fmt.Println("Number is odd")
	   } else {
			fmt.Println("Number is even")
	   }
}