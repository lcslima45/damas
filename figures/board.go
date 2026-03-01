package figures

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lcslima45/damas/colors"
)

type Board struct {
	dim  int
	size int
}

func NewBoard() *Board {
	return &Board{
		dim:  8,
		size: 80,
	}
}

func (b *Board) Draw() *fyne.Container {
	var objects []fyne.CanvasObject

	for row := 0; row < b.dim; row++ {
		for col := 0; col < b.dim; col++ {
			sq := NewSquare()
			sq.SetSize(b.size)

			if (row+col)%2 == 0 {
				sq.SetColor(colors.ColorOrange)
			} else {
				sq.SetColor(colors.ColorBlack)
			}

			obj := sq.Draw()

			obj.Move(fyne.NewPos(
				float32(col*b.size),
				float32(row*b.size),
			))

			objects = append(objects, obj)
		}
	}

	return container.NewWithoutLayout(objects...)
}
