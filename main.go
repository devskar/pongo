package main

import (
	"github.com/devskar/pongo/pongo"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(pongo.ScreenWidth, pongo.ScreenHeight)
	ebiten.SetWindowTitle("Pongo")
	g := pongo.NewGame()
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
