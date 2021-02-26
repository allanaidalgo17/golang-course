package main

import "math"

// starting with Capital Letter means PUBLIC (visible inside and outside the package)
// starting with lower case means PRIVATE (visible only inside the PACKAGE)

// example:
// Variable - is public
// variable or _Variable - is private

// Point represents a coordinate
type Point struct {
	x float64
	y float64
}

func sides(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distance calculates the distance between two points
func Distance(a, b Point) float64 {
	cx, cy := sides(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
