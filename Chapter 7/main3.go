package main
import (
 "fmt"
//  "strconv"
)
type point struct {
	x float32
	y float32
	z float32
}

func newPoint(x, y, z float32) *point {
	// Making a Copy of a Struct
	p := point{x: x, y: y, z: z}
	return &p
}

func main() {
	pt4 := newPoint(7.8, 9.1, 2.3)
	fmt.Println(pt4) // &{7.8 9.1 2.3}
	fmt.Println(pt4.x) // 7.8

	pt5 := pt4
	pt5.x = 0
	fmt.Println(pt4) // &{0 9.1 2.3}
	fmt.Println(pt5) // &{0 9.1 2.3

	pt6 := *pt4
	pt6.z = 0
	fmt.Println(pt4) // &{7.8 9.1 2.3}
	fmt.Println(pt6) // {7.8 9.1 0}

}