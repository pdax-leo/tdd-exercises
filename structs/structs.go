package structs

import "math"

// Shape struct
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	width  float64
	height float64
}

// Area func
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

// Perimeter func
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

// Circle struct
type Circle struct {
	radius float64
}

// Area func
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Smeagol struct
type Smeagol struct {
	eyeSize float64
}

// Area func
func (s Smeagol) Area() float64 {
	return s.eyeSize
}
