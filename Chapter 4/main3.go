package main
import (
 "fmt"
 "strings"
)
func main() {
	// infinite loop
	for {
		fmt.Println("Enter QUIT to exit")
		var input string
		fmt.Print("Please enter a string:")
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "QUIT" {
			break
		}
	}
}