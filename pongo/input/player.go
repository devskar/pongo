package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayerInput struct {
	Up      bool
	Down    bool
	upKey   ebiten.Key
	downKey ebiten.Key
}

func NewPlayerInput(upKey ebiten.Key, downKey ebiten.Key) PlayerInput {
	return PlayerInput{
		Up:      false,
		Down:    false,
		upKey:   upKey,
		downKey: downKey,
	}
}

func (input *PlayerInput) Update() {
	if inpututil.IsKeyJustPressed(input.upKey) {
		input.Up = true
	}

	if inpututil.IsKeyJustReleased(input.upKey) {
		input.Up = false
	}

	if inpututil.IsKeyJustPressed(input.downKey) {
		input.Down = true
	}

	if inpututil.IsKeyJustReleased(input.downKey) {
		input.Down = false
	}
}
