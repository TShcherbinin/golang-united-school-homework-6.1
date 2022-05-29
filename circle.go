package golang_united_school_homework

import (
	"fmt"
	"math"
)

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) CalcArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle (Radius: %f)", c.Radius)
}
