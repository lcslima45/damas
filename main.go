package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/lcslima45/damas/board"
)

func main() {
	a := app.New()

	w := a.NewWindow("Damas")

	board := board.NewBoard()

	w.SetContent(board.Draw())
	w.Resize(fyne.NewSize(640, 640))
	w.ShowAndRun()
}
