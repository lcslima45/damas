package board

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lcslima45/damas/colors"
	"github.com/lcslima45/damas/figures"
	"github.com/lcslima45/damas/logic"
	"github.com/lcslima45/damas/pieces"
)

type Board struct {
	bLogic  *logic.BoardLogic
	squares [8][8]*figures.Square
	size    float32
}

func NewBoard() *Board {
	return &Board{
		bLogic: logic.NewBoard(),
		size:   80,
	}
}

func (b *Board) DrawSquares() []fyne.CanvasObject {
	var objects []fyne.CanvasObject
	for row := 0; row < len(b.bLogic.Grid); row++ {
		for column := 0; column < len(b.bLogic.Grid[0]); column++ {
			sq := figures.NewSquare()
			sq.SetSize(b.size)
			if b.bLogic.Grid[row][column] == logic.Black {
				sq.SetColor(colors.ColorBlack)
			} else {
				sq.SetColor(colors.ColorOrange)
			}
			cell := sq.Draw()
			cell.Move(fyne.NewPos(
				float32(column)*b.size,
				float32(row)*b.size,
			))
			b.squares[row][column] = sq
			objects = append(objects, cell)
		}
	}
	return objects
}

func (b *Board) DrawPieces() []fyne.CanvasObject {
	var objects []fyne.CanvasObject
	for row := 0; row < len(b.bLogic.Pieces); row++ {
		for column := 0; column < len(b.bLogic.Pieces[0]); column++ {
			switch b.bLogic.Pieces[row][column] {
			case logic.Blue:
				piece := pieces.NewPiece(0.8*b.size, colors.ColorBlue)
				piece.Row = row
				piece.Column = column
				piece.OnMoveRequest = b.OnPieceMoveRequest
				posPiece := b.CalcPositionPiece(row, column)
				piece.Move(posPiece)
				objects = append(objects, piece)
			case logic.Brown:
				piece := pieces.NewPiece(0.8*b.size, colors.ColorBrown)
				piece.Row = row
				piece.Column = column
				piece.OnMoveRequest = b.OnPieceMoveRequest
				posPiece := b.CalcPositionPiece(row, column)
				piece.Move(posPiece)
				objects = append(objects, piece)
			}
		}
	}
	return objects
}

func (b *Board) Draw() *fyne.Container {
	sqObj := b.DrawSquares()
	piObj := b.DrawPieces()
	allObj := append(sqObj, piObj...)
	return container.NewWithoutLayout(allObj...)
}

func (b *Board) CalcPositionPiece(row, col int) fyne.Position {
	pieceSize := 0.8 * b.size
	offset := (b.size - pieceSize) / 2

	x := float32(col)*b.size + offset
	y := float32(row)*b.size + offset

	return fyne.NewPos(x, y)
}

func (b *Board) OnPieceMoveRequest(row, col int) {
	previews := b.bLogic.PreviewPositions(row, col)
	for _, p := range previews {
		b.squares[p.Row][p.Column].SetColor(colors.ColorBlackSelected)
	}
}
