package golang_united_school_homework

import (
	"fmt"
	"math"
)

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	return t.Side * 3
}

func (t Triangle) CalcArea() float64 {
	return math.Sqrt(3) / 4 * t.Side * t.Side
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle (Side: %f)", t.Side)
}
