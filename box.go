package golang_united_school_homework

import "fmt"

const errorFmtLiteral string = "Invalid index: %d"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

func (b *box) Size() int {
	return len(b.shapes)
}

func (b *box) Capacity() int {
	return b.shapesCapacity
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return fmt.Errorf("Exceeded box capacity")
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf(errorFmtLiteral, i)
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf(errorFmtLiteral, i)
	}
	var shape Shape = b.shapes[i]

	// remove element by inde with correct ordering
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf(errorFmtLiteral, i)
	}

	var oldShape Shape = b.shapes[i]

	b.shapes[i] = shape

	return oldShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var result float64 = 0.0

	for _, shape := range b.shapes {
		result += shape.CalcPerimeter()
	}

	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var result float64 = 0.0

	for _, shape := range b.shapes {
		result += shape.CalcArea()
	}

	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	result := []Shape{}
	foundCircles := false

	for _, shape := range b.shapes {
		switch shape.(type) {
		case Circle, *Circle:
			foundCircles = true
		default:
			result = append(result, shape)
		}
	}

	if foundCircles {
		b.shapes = result
		return nil
	}

	return fmt.Errorf("Circles not found")
}
