package figures

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lcslima45/damas/colors"
)

type Board struct {
	dim  int
	size float32
}

func NewBoard() *Board {
	return &Board{
		dim:  8,
		size: 80,
	}
}

func (b *Board) Draw() *fyne.Container {
	var objects []fyne.CanvasObject

	//Draw squares for cells in board
	for row := 0; row < b.dim; row++ {
		for col := 0; col < b.dim; col++ {
			sq := NewSquare()
			sq.SetSize(b.size)

			if (row+col)%2 == 0 {
				sq.SetColor(colors.ColorOrange)
			} else {
				sq.SetColor(colors.ColorBlack)
			}

			cell := sq.Draw()

			cell.Move(fyne.NewPos(
				float32(col)*b.size,
				float32(row)*b.size,
			))

			objects = append(objects, cell)
		}
	}

	//draw circles for pieces inside cells of board
	for row := 0; row < b.dim; row++ {
		if row >= 3 && row < 5 {
			continue
		}
		for col := 0; col < b.dim; col++ {
			circle := NewCircle()
			circle.SetSize(0.8 * b.size)
			if row <= 2 && (row+col)%2 != 0 {
				circle.SetColor(colors.ColorBlue)
			} else if row >= 5 && (row+col)%2 != 0 {
				circle.SetColor(colors.ColorBrown)
			} else {
				continue
			}

			piece := circle.Draw()

			posPiece := b.CalcPositionPiece(row, col)
			piece.Move(posPiece)
			objects = append(objects, piece)
		}
	}
	return container.NewWithoutLayout(objects...)
}

func (b *Board) CalcPositionPiece(row, col int) fyne.Position {
	pieceSize := 0.8 * b.size
	offset := (b.size - pieceSize) / 2

	x := float32(col)*b.size + offset
	y := float32(row)*b.size + offset

	return fyne.NewPos(x, y)
}
