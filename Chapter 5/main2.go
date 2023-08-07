package main
import (
	"fmt"
	"time"
   )
func main() {
	// Defining functions with parameters
	displayDate("Mon 2006-01-02 15:04:05")
}

func displayDate(format string) {
	fmt.Println(time.Now().Format(format))
}