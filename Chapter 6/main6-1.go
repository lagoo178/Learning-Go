package main

import "fmt"

func main() {
    // Create a source slice.
    source := []int{1, 2, 3, 4, 5}

    // Create a destination slice with enough capacity to hold all of the elements from the source slice.
    destination := make([]int, len(source))

    // Copy the elements from the source slice to the destination slice.
    copy(destination, source)

    // Print the destination slice.
    fmt.Println(destination) // Output: [1 2 3 4 5]
}