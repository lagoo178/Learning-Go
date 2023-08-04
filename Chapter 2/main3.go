package main

import (
	"fmt"
	// "strconv"
   )

// var num1 = 5 // type inferred

func main() {
	queue := 5
	name := "John"
	// s := name + ", your queue number is:" +
	// strconv.Itoa(queue)

	s := fmt.Sprintf("%s, your queue number is %d",
 		name, queue)

    fmt.Println(s)
}