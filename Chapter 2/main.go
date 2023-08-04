package main

import (
	"fmt"
	"time"
   )

// var num1 = 5 // type inferred

func main() {
	var num1 = 5 // type inferred
	var num2 int = 25// explicitly typed
	var num3 float32 = 4.5// floating point variable
	var name string = "Ayam Geprek"// string variable
	var raining bool = true // boolean variable
	var firstName, lastName, age = "Wei-Meng", "Lee", 25
	jawa := "sampun mas"
	_ = "sampun mas"
	start := time.Now()

	fmt.Println(num1) // 5
	fmt.Println(num2) // default = 0
	fmt.Println(num3) // default = 0
	fmt.Println(name) // default = "" (empty string)
	fmt.Println(raining) // default = false
	fmt.Println(firstName, lastName, age)
	fmt.Println(jawa)
	fmt.Printf("%T\n", start)
}