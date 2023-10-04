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
	// Creating a Go Struct
	p := point{x: x, y: y, z: z}
	return &p
}

func main() {
	pt4 := newPoint(7.8, 9.1, 2.3)
	fmt.Println(pt4) // &{7.8 9.1 2.3}
	fmt.Println(pt4.x) // 7.8

}