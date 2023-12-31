package main

import (
	"fmt"
	"math"
)

type Shape interface {
 	Area() float64
}

type Circle struct {
 	radius float64
}

type Square struct {
 	length float64
}

type Triangle struct {
	base float64
	height float64
}

// Circle implements Shape
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}


func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

// Square implements Shape
func (s Square) Area() float64 {
 	return s.length * s.length
}

// Triangle implements Shape
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func calculateArea(listOfShapes []Shape) {
	for _, shape := range listOfShapes {
		fmt.Println("Area: ", shape.Area())
	}
	
}
func main() {
	c1 := Circle{radius: 5}
	s1 := Square{length: 6}
	t1 := Triangle{base: 6, height: 8}
	shapes := []Shape{c1, s1, t1}
	calculateArea(shapes)
}