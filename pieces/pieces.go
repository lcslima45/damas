package pieces

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/lcslima45/damas/colors"
	"github.com/lcslima45/damas/figures"
)

type Piece struct {
	widget.BaseWidget
	Circle    *figures.Circle
	posRow    int
	posColumn int
}

func NewPiece(size float32, col color.Color) *Piece {
	c := figures.NewCircle()

	c.SetSize(size)
	c.SetColor(col)

	p := &Piece{
		Circle: c,
	}

	p.Resize(fyne.NewSize(size, size))
	p.ExtendBaseWidget(p)
	return p
}

func (p *Piece) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(p.Circle.Draw())
}

func (p *Piece) Tapped(*fyne.PointEvent) {
	switch p.Circle.GetColor() {
	case colors.ColorBlue:
		p.Circle.SetColor(colors.ColorGreen)
		// p.CalcNewPostion()
	case colors.ColorGreen:
		p.Circle.SetColor(colors.ColorBlue)
	case colors.ColorBrown:
		p.Circle.SetColor(colors.ColorRed)
		// p.CalcNewPostion()
	case colors.ColorRed:
		p.Circle.SetColor(colors.ColorBrown)
	}

	p.Refresh()
}

func (p *Piece) CalcNewPostion() {

}
