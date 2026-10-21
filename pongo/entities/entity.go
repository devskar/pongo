package entities

import (
	"github.com/devskar/pongo/pongo/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Draw(screen *ebiten.Image)
}

type PlayerEntity interface {
	Update(input *input.PlayerInput)
}
