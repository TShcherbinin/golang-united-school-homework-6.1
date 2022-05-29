package golang_united_school_homework

import "fmt"

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcPerimeter() float64 {
	return (r.Weight + r.Height) * 2
}

func (r Rectangle) CalcArea() float64 {
	return r.Weight * r.Height
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle (Height: %f; Weight: %f)", r.Height, r.Weight)
}
