package main
import (
 "fmt"
//  "strings"
)
func main() {
	//iterate array through string
	for pos, char := range "Hello, world!" {
		fmt.Printf("%d %c\n", pos, char)
	}
	   
}