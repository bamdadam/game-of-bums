package internal

import (
	"errors"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	gamestate   *GameState
	gameWidth   int
	pixMapWidth int
	pixColor    []byte
	tick        uint8
}

type GameState struct {
	pixels          []byte
	pixmaps         []Pixmap
	state           []bool
	gameStateLength int
	tickCounter     uint8
}

func NewGameState(gameSize, pixmapLen int, startingState []bool) (*GameState, error) {
	if len(startingState)*int(math.Pow(float64(pixmapLen), 2)) != int(math.Pow(float64(gameSize), 2)) {
		return nil, errors.New("game size !=  starting state * pixmap length")
	}
	if gameSize%pixmapLen != 0 {
		return nil, errors.New("game length should be divisible by pixmap length")
	}
	baseGameState := make([]bool, int(math.Pow(float64(gameSize), 2)))
	pixmaps := []Pixmap{}
	for i := 0; i < len(startingState); i++ {
		pixmap := NewPixmap(i, pixmapLen, startingState[i])
		pixmap.InitPixmap(gameSize, baseGameState)
		pixmap.FillPixmap()
		pixmaps = append(pixmaps, *pixmap)
	}
	return &GameState{
		pixels:          make([]byte, 4*gameSize*gameSize),
		state:           baseGameState,
		pixmaps:         pixmaps,
		gameStateLength: gameSize,
		tickCounter:     0,
	}, nil
}

func NewGame(gameSize, pixmapSize int, startingState []bool, iterPerTick uint8) (*Game, error) {
	ss, err := NewGameState(gameSize, pixmapSize, startingState)
	if err != nil {
		return nil, errors.New("error while creating new game state: " + err.Error())
	}
	return &Game{
		gamestate:   ss,
		gameWidth:   gameSize,
		pixMapWidth: pixmapSize,
		// Green
		pixColor: []byte{0, 255, 0, 255},
		tick:     iterPerTick,
	}, nil
}

func (g *Game) Update() error {
	if !g.canRun(g.tick) {
		g.gamestate.tickCounter++
		return nil
	}
	g.runGeneration()
	g.fillGamePixmaps()
	for index, v := range g.gamestate.state {
		if v {
			copy(g.gamestate.pixels[4*index:4*index+4], g.pixColor)
		} else {
			copy(g.gamestate.pixels[4*index:4*index+4], []byte{0, 0, 0, 255})
		}
	}
	g.gamestate.tickCounter++
	return nil
}

// Calculate the new generation based on the old generation and the game rules
func (g *Game) runGeneration() {
	newGamePixmap := make([]Pixmap, len(g.gamestate.pixmaps))
	for index, v := range g.gamestate.pixmaps {
		adjs := generateSliceFromIndex(g.gamestate.pixmaps, getNeighbors(index, g.gameWidth/g.pixMapWidth)...)
		c := getAdjAlives(adjs)
		if v.IsAlive() && (c == 2 || c == 3) {
			newGamePixmap[index] = g.gamestate.pixmaps[index]
			newGamePixmap[index].SetAlive(true)
		} else if !v.IsAlive() && c == 3 {
			newGamePixmap[index] = g.gamestate.pixmaps[index]
			newGamePixmap[index].SetAlive(true)
		} else {
			newGamePixmap[index] = g.gamestate.pixmaps[index]
			newGamePixmap[index].SetAlive(false)
		}
	}
	g.gamestate.pixmaps = newGamePixmap
}

func (g *Game) fillGamePixmaps() {
	for _, v := range g.gamestate.pixmaps {
		v.FillPixmap()
	}
}

func (g *Game) canRun(tick uint8) bool {
	if g.gamestate.tickCounter%tick != 0 {
		return false
	}
	return true
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.gamestate.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
