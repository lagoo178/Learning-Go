package main
import (
 "fmt"
//  "strings"
)
func main() {
	//iterate array
	var OS [3]string
	OS[0] = "iOS"
	OS[1] = "Android"
	OS[2] = "Windows"
	for i, v := range OS {
		fmt.Println(i, v)
	
	}


}