package main

import (
	"log"
	"main/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.NewGame()
	ebiten.SetWindowSize(1440, 720)
	ebiten.SetWindowTitle("График с динамическим добавлением точек")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
