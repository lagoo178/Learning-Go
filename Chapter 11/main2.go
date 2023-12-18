package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Seeing how shared resources impact goroutines

// This program will :
// » The credit() function adds 100 to balance a total of five times.
// » The debit() function deducts 100 from balance a total of five times.

var balance int
var mutex = &sync.Mutex{}

func credit() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		balance += 100
		time.Sleep(time.Duration(rand.Intn(100)) *
		time.Millisecond)
		fmt.Println("After crediting, balance is",
		balance)
		mutex.Unlock()
	}
}
func debit() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		balance -= 100
		time.Sleep(time.Duration(rand.Intn(100)) *
		time.Millisecond)
		fmt.Println("After debiting, balance is", balance)
		mutex.Unlock()
	}
}
func main() {
	balance = 200
	fmt.Println("Initial balance is", balance)
	go credit()
	go debit()
	fmt.Scanln()
}