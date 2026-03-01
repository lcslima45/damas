package main

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/lcslima45/damas/colors"
	"github.com/lcslima45/damas/figures"
)

func main() {
	a := app.New()
	w := a.NewWindow("Square Window")

	var square figures.Figure = figures.NewSquare()
	square.SetSize(200)
	square.SetColor(colors.ColorBlack)
	obj := square.Draw()

	var circle figures.Figure = figures.NewCircle()
	circle.SetSize(200)
	circle.SetColor(colors.ColorOrange)
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
