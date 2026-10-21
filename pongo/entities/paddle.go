package entities

import (
	"image/color"

	"github.com/devskar/pongo/pongo/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddle struct {
	fill   color.Color
	Width  float32
	Height float32
	X      float32
	Y      float32
}

func NewPaddle(x float32, y float32, width float32, height float32) Paddle {
	return Paddle{
		fill:   color.White,
		Width:  width,
		Height: height,
		X:      x,
		Y:      y,
	}
}

func (t Paddle) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, t.X, t.Y, t.Width, t.Height, t.fill, false)
}

func (t *Paddle) Update(input *input.PlayerInput) {
	paddleNextY := t.Y

	if input.Up {
		paddleNextY -= 10
	}

	if input.Down {
		paddleNextY += 10
	}

	_, height := ebiten.ScreenSize()

	if 0 > paddleNextY || paddleNextY+t.Height > float32(height) {
		return
	}
	t.Y = paddleNextY
}
