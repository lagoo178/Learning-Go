package main
import (
	"fmt"
	"math"
)
type point struct {
	x float32
	y float32
	z float32
}
func (p point) length() float64 {
	// Defining Methods in Structs
	return math.Sqrt(
		(math.Pow(float64(p.x), 2) +
		math.Pow(float64(p.y), 2) +
		math.Pow(float64(p.z), 2)))
}
func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}

func main() {
	pt4 := newPoint(7.8, 9.1, 2.3)
	fmt.Println(pt4.length()) // 12.2040980698644
}