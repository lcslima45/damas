package figures

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type Figure interface {
	SetSize(size float32)
	SetColor(c color.Color)
	Draw() fyne.CanvasObject
}
