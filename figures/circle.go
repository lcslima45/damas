package figures

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Circle struct {
	size   float32
	color  color.Color
	object *canvas.Circle
}

func NewCircle() *Circle {
	return &Circle{
		object: canvas.NewCircle(color.Transparent),
	}
}

func (c *Circle) SetSize(size float32) {
	c.size = size
	c.object.Resize(fyne.NewSize(size, size))
}

func (c *Circle) SetColor(color color.Color) {
	c.color = color
	c.object.FillColor = color
	c.object.Refresh()
}

func (c *Circle) GetColor() color.Color {
	return c.color
}

func (c *Circle) Draw() fyne.CanvasObject {
	return c.object
}
