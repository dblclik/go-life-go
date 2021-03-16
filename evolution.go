package main

import "sync"

/*
These rules, which compare the behavior of the automaton to real life,
can be condensed into the following:

	- Any live cell with two or three live neighbours survives.
	- Any dead cell with three live neighbours becomes a live cell.
	- All other live cells die in the next generation. Similarly, all other dead cells stay dead.
*/

// Evolve iterates the matrix forward by 1 step
func Evolve(x int, y int, matIn [][]int, matOut [][]int, wg *sync.WaitGroup) {
	// Start with an active count of 0
	defer wg.Done()

	activeCells := 0
	yMax := len(matIn)
	xMax := len(matIn[0])

	neighborhood := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, cells := range neighborhood {
		xIndex := (x + cells[1] + xMax) % xMax
		yIndex := (y + cells[0] + yMax) % yMax
		activeCells += matIn[yIndex][xIndex]
	}

	if activeCells > 3 || activeCells < 2 {
		matOut[y][x] = 0
	} else {
		if activeCells == 3 {
			matOut[y][x] = 1
		} else {
			matOut[y][x] = 1 * matIn[y][x]
		}
	}
}
