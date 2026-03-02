package figures

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Circle struct {
	size  float32
	color color.Color
}

func NewCircle() Figure {
	return &Circle{}
}

func (c *Circle) SetSize(size float32) {
	c.size = size
}

func (c *Circle) SetColor(color color.Color) {
	c.color = color
}

func (c *Circle) Draw() fyne.CanvasObject {
	circle := canvas.NewCircle(c.color)

	circle.Resize(fyne.NewSize(c.size, c.size))
	return circle
}
