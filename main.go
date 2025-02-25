package main

import (
	"log"
	"main/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &game.Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("График с динамическим добавлением точек")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
