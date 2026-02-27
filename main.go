package main

import (
	"image/color"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var (
	ColorBlack  = color.Black
	ColorOrange = color.RGBA{255, 165, 0, 255}
)

type Figure interface {
	SetSize(size int)
	SetColor(c color.Color)
	Draw() fyne.CanvasObject
}

type Square struct {
	size  int
	color color.Color
}

type Circle struct {
	size  int
	color color.Color
}

func NewSquare() Figure {
	return &Square{}
}

func (s *Square) SetSize(size int) {
	s.size = size
}

func (s *Square) SetColor(c color.Color) {
	s.color = c
}

func (s *Square) Draw() fyne.CanvasObject {
	square := canvas.NewRectangle(s.color)

	square.SetMinSize(fyne.NewSize(float32(s.size), float32(s.size)))
	return square
}

func NewCircle() Figure {
	return &Circle{}
}

func (c *Circle) SetSize(size int) {
	c.size = size
}

func (c *Circle) SetColor(color color.Color) {
	c.color = color
}

func (c *Circle) Draw() fyne.CanvasObject {
	circle := canvas.NewCircle(c.color)

	circle.Resize(fyne.NewSize(float32(c.size), float32(c.size)))
	return circle
}

func main() {
	a := app.New()
	w := a.NewWindow("Square Window")

	var square Figure = NewSquare()
	square.SetSize(200)
	square.SetColor(ColorBlack)
	obj := square.Draw()

	var circle Figure = NewCircle()
	circle.SetSize(200)
	circle.SetColor(ColorOrange)
	obj2 := circle.Draw()

	// MUITO IMPORTANTE
	obj.Resize(fyne.NewSize(200, 200))
	obj2.Resize(fyne.NewSize(200, 200))

	obj.Move(fyne.NewPos(0, 200))
	obj2.Move(fyne.NewPos(220, 0))

	content := container.NewWithoutLayout(obj, obj2)

	w.Resize(fyne.NewSize(800, 800))
	w.SetContent(content)
	w.ShowAndRun()
}
