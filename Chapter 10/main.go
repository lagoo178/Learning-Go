package main

import (
 	"fmt"
)

// Implementing an interface

type DigitsCounter interface {
	CountOddEven() (int, int)
}

type DigitString string
// DigitString implements DigitsCounter
func (ds DigitString) CountOddEven() (int, int) {
	odds, evens := 0, 0
	for _, digit := range ds {
		if digit%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}
func main() {
	s := DigitString("123456789")
	var d DigitsCounter
	d = s
	fmt.Println(d.CountOddEven()) // 5 4
}