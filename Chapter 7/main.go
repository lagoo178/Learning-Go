package main
import (
 "fmt"
//  "strconv"
)

type point struct {
	// Defining Structs for a Collection of Items
	x float32
	y float32
	z float32
}

func main() {
	var pt1 point
	pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 4.2

	pt2 := point{
					x: 5.6, 
					y: 3.8, 
					z: 6.9,
				}

	pt3 := point{x: 5.6, y: 3.8}
	
	fmt.Println(pt1.x)
	fmt.Println(pt1.y)
	fmt.Println(pt1.z)
	fmt.Println(pt2) // {5.6 3.8 6.9}
	fmt.Println(pt3) // {5.6 3.8 0}

}