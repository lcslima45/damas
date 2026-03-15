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

type Position struct {
	Row    int
	Column int
}

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

func (b *BoardLogic) PreviewPositions(row, column int) []Position {
	if row >= 0 && row <= 7 && column >= 0 && column <= 7 {
		if b.Pieces[row][column] == Blue {
			if column >= 1 && column <= 6 {
				Pos1 := Position{
					Row:    row + 1,
					Column: column + 1,
				}
				Pos2 := Position{
					Row:    row + 1,
					Column: column - 1,
				}
				return []Position{Pos1, Pos2}
			} else if column == 0 {
				Pos1 := Position{
					Row:    row + 1,
					Column: column + 1,
				}
				return []Position{Pos1}
			} else if column == 7 {
				Pos1 := Position{
					Row:    row + 1,
					Column: column - 1,
				}
				return []Position{Pos1}
			}
		}

		if b.Pieces[row][column] == Brown {
			if column >= 1 && column <= 6 {
				Pos1 := Position{
					Row:    row - 1,
					Column: column + 1,
				}
				Pos2 := Position{
					Row:    row - 1,
					Column: column - 1,
				}
				return []Position{Pos1, Pos2}
			} else if column == 0 {
				Pos1 := Position{
					Row:    row - 1,
					Column: column + 1,
				}
				return []Position{Pos1}
			} else if column == 7 {
				Pos1 := Position{
					Row:    row - 1,
					Column: column - 1,
				}

				return []Position{Pos1}
			}
		}
	}
	return nil
}
