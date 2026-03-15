package logic

type BoardLogic struct {
	Grid   [8][8]SquareLogic
	Pieces [8][8]PiecesLogic
}

type SquareLogic int

type PiecesLogic int

const (
	Orange SquareLogic = iota
	Black
)

const (
	Blue PiecesLogic = iota + 1
	Brown
)

func NewBoard() *BoardLogic {
	b := &BoardLogic{}
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if (row+column)%2 == 0 {
				b.Grid[row][column] = Black
			} else {
				b.Grid[row][column] = Orange
			}
		}
	}

	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if b.Grid[row][column] == Black {
				switch row {
				case 0, 1:
					b.Pieces[row][column] = Blue
				case 6, 7:
					b.Pieces[row][column] = Brown
				}
			}
		}
	}
	return b
}
