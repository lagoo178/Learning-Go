package main

import (
	"fmt"
	"strconv"
   )

// var num1 = 5 // type inferred

func main() {
	var input string
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%s", &input)
	age, err := strconv.Atoi(input) // convert string to
 									// int
	if err != nil { // an error occurred
		fmt.Println(err)
	} else {
		fmt.Println("Your age is:", age)
	}
}