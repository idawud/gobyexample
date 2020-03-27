package structs

type Rectangle struct {
	Height float64
	Width float64
}

func (rectangle Rectangle) Perimeter() float64 {
    return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle Rectangle) Area() float64 {
    return rectangle.Width * rectangle.Height
}