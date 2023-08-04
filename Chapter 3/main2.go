package main

import (
	"fmt"
	// "strconv"
   )

func main() {
	if !raining() || snowing() {
		fmt.Println("Let's ski!")
	}
	   /*
	   Check if it is raining now...
	   Check if it is snowing now...
	   Let's ski!
	   */
	   
}

func raining() bool {
	fmt.Println("Check if it is raining now...")
	return true
}
func snowing() bool {
	fmt.Println("Check if it is snowing now...")
	return true
}
   