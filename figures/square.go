package figures

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Square struct {
	size   float32
	color  color.Color
	object *canvas.Rectangle
}

func NewSquare() *Square {
	return &Square{
		object: canvas.NewRectangle(color.Transparent),
	}
}

func (s *Square) SetSize(size float32) {
	s.size = size
	s.object.Resize(fyne.NewSize(size, size))
}

func (s *Square) SetColor(c color.Color) {
	s.color = c
	s.object.FillColor = s.color
	s.object.Refresh()
}

func (s *Square) Draw() fyne.CanvasObject {
	return s.object
}
