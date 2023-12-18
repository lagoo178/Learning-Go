package main
import (
	"fmt"
	"time"
)

// Understanding Channels

//---send data into a channel---
func sendData(ch chan string) {
	fmt.Println("Sending a string into channel...")
	// comment out the following line
	// time.Sleep(2 * time.Second)
	
	ch <- "Hello"
	fmt.Println("String has been retrieved from channel...")
}

//---getting data from the channel---
func getData(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("String retrieved from channel:", <-ch)
}

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	fmt.Scanln()
}