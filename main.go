package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {

	const (
		screenWidth  = 640
		screenHeight = 480
	)
	game := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("THE GAME OF BUMS")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal("Panic while running game: ", err)
	}
}
