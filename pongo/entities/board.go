package entities

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	borderElementCount = 9
	borderElementWidth = 10
)

var (
	col = color.RGBA{120, 120, 120, 255}
)

type Board struct {
}

func (b Board) Draw(screen *ebiten.Image) {
	drawBoarder(screen)
	drawCenter(screen)
}

func drawCenter(screen *ebiten.Image) {
	size := screen.Bounds().Size()

	vector.FillCircle(screen, float32(size.X)/2, float32(size.Y)/2, 30, col, false)
	vector.FillCircle(screen, float32(size.X)/2, float32(size.Y)/2, 25, color.Black, false)
}

func drawBoarder(screen *ebiten.Image) {
	size := screen.Bounds().Size()

	partLen := size.Y / (borderElementCount*2 - 1)

	vector.FillRect(screen, 0, 0, float32(size.X), borderElementWidth, col, false)
	vector.FillRect(screen, 0, float32(size.Y)-borderElementWidth, float32(size.X), borderElementWidth, col, false)

	for x := range borderElementCount {
		offset := 2 * x * partLen

		vector.FillRect(screen, 0, float32(offset), borderElementWidth, float32(partLen), col, false)
		vector.FillRect(screen, float32(size.X)-borderElementWidth, float32(offset), borderElementWidth, float32(partLen), col, false)
	}

}
