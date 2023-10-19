# Game-Of-Bums
Conway's game of life implemented in golang using `[Ebitengine](https://github.com/hajimehoshi/ebiten)`

## Introduction

The Game of Life, also known simply as Life, is a cellular automaton devised by John Horton Conway in 1970. It is a zero-player game, meaning that its evolution is determined by its initial state. One interacts with the Game of Life by creating an initial configuration and observing how it evolves. It is Turing complete and can simulate a universal constructor or any other Turing machine.

## Structure
### Pixmap
- A Square of length N grouping of pixels (meaning each Pixmap holds N*N pixels)
- Each Pixmap has a reletive index based on its position in the 2D game window
```go
type Pixmap struct {
	pixels []*bool
	idx    int
	size   int
	status bool
}
```
### Gamestate
- State holds the state (being alive or dead) of the whole board
- In each iteration of the game, a new state is created from the current state and the game's rules
- Based on the created game state, Pixels slice is filled which is then shown on the screen using ebitengine
```go
type GameState struct {
	pixels          []byte
	pixmaps         []Pixmap
	state           []bool
	gameStateLength int
	tickCounter     uint8
}
```
## Run
- To run the code simply set the configs based on your liking
- Iter per tick determines the speed for which each generation runs
```go
	const (
		screenWidth = 600
		pixmapLen   = 5
		iterPerTick = 8
	)
```
 - Run the simulation: `make run`