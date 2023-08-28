package main
import (
 "fmt"
//  "strconv"
)
func main() {
 	// Making copies of an array or slice

	nums1 := [5]int{1, 2, 3, 4, 5}
	nums2 := nums1
	nums2[0] = 77 // If you wanna change nums2, the changes should only affect nums2 and not nums1
	fmt.Println(nums1) // 1 2 3 4 5]
	fmt.Println(nums2) // 77 2 3 4 5
}