package main
import (
	"fmt"
	"time"
   )
func main() {
	// Defining functions with multiple parameters
	displayDate("Mon 2006-01-02 15:04:05", "Current Date and Time:")
}

func displayDate(format string, prefix string) {
	fmt.Println(prefix, time.Now().Format(format))
   }