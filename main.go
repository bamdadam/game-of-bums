package main

import (
	"crypto/rand"
	"errors"
	"log"
	"math"
	"math/big"

	"github.com/bamdadam/game-of-bums/internal"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameState struct {
	pixels  []byte
	kernels []internal.Kernel
	state   []bool
	height  int
	counter uint8
}

func NewState(gameSize, kernelLen int, startingState []bool) (*GameState, error) {
	if len(startingState)*int(math.Pow(float64(kernelLen), 2)) != int(math.Pow(float64(gameSize), 2)) {
		return nil, errors.New("game size !=  starting state * kernel len")
	}
	if gameSize%kernelLen != 0 {
		return nil, errors.New("game len should be divisible by kernel len")
	}
	baseGameState := make([]bool, int(math.Pow(float64(gameSize), 2)))
	kernels := []internal.Kernel{}
	for i := 0; i < len(startingState); i++ {
		kernel := internal.NewKernel(i, kernelLen, startingState[i])
		kernel.InitKernel(gameSize, baseGameState)
		kernel.FillKernel()
		kernels = append(kernels, *kernel)

	}
	return &GameState{
		pixels:  make([]byte, 4*gameSize*gameSize),
		state:   baseGameState,
		kernels: kernels,
		height:  gameSize,
		counter: 0,
	}, nil
}

type Game struct {
	gamestate   *GameState
	gameWidth   int
	kernelWidth int
}

func NewGame(gameSize, kernelSize int, startingState []bool) (*Game, error) {
	ss, err := NewState(gameSize, kernelSize, startingState)
	if err != nil {
		return nil, errors.New("error while creating new state: " + err.Error())
	}
	return &Game{
		gamestate:   ss,
		gameWidth:   gameSize,
		kernelWidth: kernelSize,
	}, nil
}

func (g *Game) Update() error {
	if !g.canRun() {
		g.gamestate.counter++
		return nil
	}
	fillerColor := g.calcColor()
	g.runGeneration()
	// fmt.Println(g.gamestate.state)
	// for _, v := range g.gamestate.kernels {
	// 	fmt.Println(v.String())
	// }
	// fmt.Println("========")
	g.fillGameKernels()
	for index, v := range g.gamestate.state {
		if v {
			copy(g.gamestate.pixels[4*index:4*index+4], fillerColor)
		} else {
			copy(g.gamestate.pixels[4*index:4*index+4], []byte{0, 0, 0, 255})
		}
	}
	g.gamestate.counter++
	return nil
}

func (g *Game) runGeneration() {
	newGameKernels := make([]internal.Kernel, len(g.gamestate.kernels))
	for index, v := range g.gamestate.kernels {
		adjs := generateSliceFromIndex(g.gamestate.kernels, getNeighbors(index, g.gameWidth/g.kernelWidth)...)
		// fmt.Println(index, adjs)
		c := getAdjAlives(adjs)
		// fmt.Println(index, c, v.IsAlive())
		if v.IsAlive() && (c == 2 || c == 3) {
			newGameKernels[index] = g.gamestate.kernels[index]
			newGameKernels[index].SetAlive(true)
		} else if !v.IsAlive() && c == 3 {
			newGameKernels[index] = g.gamestate.kernels[index]
			newGameKernels[index].SetAlive(true)
		} else {
			newGameKernels[index] = g.gamestate.kernels[index]
			newGameKernels[index].SetAlive(false)
		}
	}
	g.gamestate.kernels = newGameKernels
}

func (g *Game) fillGameKernels() {
	for _, v := range g.gamestate.kernels {
		v.FillKernel()
	}
}

func getNeighbors(idx int, n int) []int {
	// Convert 1D index to 2D indices
	i := idx / n
	j := idx % n

	// Define the offsets for the 8 neighboring positions
	offsets := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	neighbors := []int{}

	for _, offset := range offsets {
		// Calculate the 2D indices of the neighbor
		ni := i + offset[0]
		nj := j + offset[1]

		// Check if the neighbor is within bounds
		if ni >= 0 && ni < n && nj >= 0 && nj < n {
			// Convert the 2D indices back to a 1D index
			neighborIdx := ni*n + nj
			neighbors = append(neighbors, neighborIdx)
		}
	}

	return neighbors
}

func generateSliceFromIndex[T any](sl []T, in ...int) []T {
	slSize := len(sl)
	res := []T{}
	for _, v := range in {
		if v >= 0 && v < slSize {
			res = append(res, sl[v])
		}
	}
	return res
}

func getAdjAlives(adj []internal.Kernel) int {
	counter := 0
	for _, v := range adj {
		if v.IsAlive() {
			counter++
		}
	}
	return counter
}

func (g *Game) canRun() bool {
	if g.gamestate.counter%8 != 0 {
		return false
	}
	return true
}

func (g *Game) calcColor() []byte {
	greenColor := []byte{0, 255, 0, 255}
	return greenColor
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.WritePixels(g.gamestate.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {

	const (
		screenWidth = 600
		kernelLen   = 6
	)
	startingState := make([]bool, int(math.Pow(screenWidth/kernelLen, 2)))
	for index := range startingState {
		rn, err := rand.Int(rand.Reader, big.NewInt(6))
		if err != nil {
			log.Fatal(err)
		}
		if rn.Int64() == 1 {
			startingState[index] = true
		}
	}
	// fmt.Println(screenWidth*screenWidth, kernelLen*kernelLen*len(startingState))
	// fmt.Println(startingState)
	game, err := NewGame(screenWidth, kernelLen, startingState)
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(screenWidth, screenWidth)
	ebiten.SetWindowTitle("THE GAME OF BUMS")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal("Panic while running game: ", err)
	}
}
