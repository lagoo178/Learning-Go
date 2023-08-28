package main
import (
 "fmt"
//  "strconv"
)

func insert(orig []int, index int, value int) ([]int, error) {
	// Inserting an item into a slice
	
	// if index < 0 {
	// 	return nil, errors.New(
	// 	"Index cannot be less than 0")
	// }
	if index >= len(orig) {
		return append(orig, value), nil
	}
	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value
	return orig, nil
}

func delete(orig []int, index int) ([]int, error) {
	// Removing an item from a slice

	// if index < 0 || index >= len(orig) {
	// 	return nil, errors.New("Index out of range")
	// }
	orig = append(orig[:index], orig[index+1:]...)
	return orig, nil
}

func main() {
	// Testing the func
	t := []int{1, 2, 3, 4, 5}
	t, err := delete(t, 2)
	// t, err := insert(t, 2, 9)
	if err == nil {
	 fmt.Println(t) // 1 2 9 3 4 5]
	} else {
	 fmt.Println(err)
	}

}