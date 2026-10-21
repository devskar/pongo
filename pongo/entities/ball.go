package entities

import (
	"image/color"

	"github.com/devskar/pongo/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	fill   color.Color
	Radius float32
	X      float32
	Y      float32
	Speed  vectors.Vector2
}

func NewBall(x float32, y float32, radius float32, speed vectors.Vector2) Ball {
	return Ball{
		fill:   color.White,
		Radius: radius,
		X:      x,
		Y:      y,
		Speed:  speed,
	}
}

func (b Ball) Draw(screen *ebiten.Image) {
	vector.FillCircle(screen, b.X, b.Y, b.Radius, b.fill, false)
}

func (b *Ball) Update() {
	b.X += b.Speed.X
	b.Y += b.Speed.Y
}
