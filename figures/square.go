package figures

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Square struct {
	size  float32
	color color.Color
}

func NewSquare() Figure {
	return &Square{}
}

func (s *Square) SetSize(size float32) {
	s.size = size
}

func (s *Square) SetColor(c color.Color) {
	s.color = c
}

func (s *Square) Draw() fyne.CanvasObject {
	square := canvas.NewRectangle(s.color)

	square.Resize(fyne.NewSize(float32(s.size), float32(s.size)))
	return square
}
