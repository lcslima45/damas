package figures

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type Figure interface {
	SetSize(size int)
	SetColor(c color.Color)
	Draw() fyne.CanvasObject
}
