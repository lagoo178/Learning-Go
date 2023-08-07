package main
import (
	"fmt"
	"time"
   )
func main() {
	displayDate()
}

func displayDate() {
	fmt.Println(time.Now().Date())
}