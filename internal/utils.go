package internal

import (
	"crypto/rand"
	"log"
	"math/big"
)

// Calculate neighbors of node at index in a 2D array of length n
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

// Make a new slice from indexes of another slice
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

// Calculate the number of alive neighbors of a node
func getAdjAlives(adj []Pixmap) int {
	counter := 0
	for _, v := range adj {
		if v.IsAlive() {
			counter++
		}
	}
	return counter
}

func GenerateStartingState(stateLen int) []bool {
	startingState := make([]bool, stateLen)
	for index := range startingState {
		rn, err := rand.Int(rand.Reader, big.NewInt(6))
		if err != nil {
			log.Fatal(err)
		}
		if rn.Int64() == 1 {
			startingState[index] = true
		}
	}
	return startingState
}
