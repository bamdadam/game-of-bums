package main

import (
	"log"
	"math"

	"github.com/bamdadam/game-of-bums/internal"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	const (
		screenWidth = 600
		pixmapLen   = 8
		iterPerTick = 8
	)
	startingState := internal.GenerateStartingState(int(math.Pow(screenWidth/pixmapLen, 2)))
	game, err := internal.NewGame(screenWidth, pixmapLen, startingState, iterPerTick)
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(screenWidth, screenWidth)
	ebiten.SetWindowTitle("THE GAME OF BUMS")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal("Panic while running game: ", err)
	}
}
