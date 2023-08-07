package main
import (
 "fmt"
//  "strings"
)
func main() {
	//Using Labels with the for Loop
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
		fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("-----------")
	   }
}