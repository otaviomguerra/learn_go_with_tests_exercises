package structs_and_methods

import "math"

type Rectangle struct {
	Height, Width float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

type Shape interface {
	Area() float64
}

//Perimeter calculates the perimeter for a Rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

//Area calculates the area of a Rectangle.
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//Perimeter calculates the perimeter for a Circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

//Area calculates the area of a Circle.
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

//Area calculates the area of a Triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
