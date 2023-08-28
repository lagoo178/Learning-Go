package main
import (
 "fmt"
//  "strconv"
)
func main() {
 	// Extracting part of an array or slice
	// var c [3]string
	// c[0] = "iOS"
	// c[1] = "Android"
	// c[2] = "Windows"
	 
	// fmt.Println(c[0:2]) // [iOS Android]

	t := []int{1, 2, 3, 4, 5}
	t = t[2:4]
	fmt.Println(t) // [3 4]
	fmt.Println(len(t)) // 2
	fmt.Println(cap(t)) // 3
}