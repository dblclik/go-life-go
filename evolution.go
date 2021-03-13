package golife

import "sync"

/*
These rules, which compare the behavior of the automaton to real life,
can be condensed into the following:

	- Any live cell with two or three live neighbours survives.
	- Any dead cell with three live neighbours becomes a live cell.
	- All other live cells die in the next generation. Similarly, all other dead cells stay dead.
*/

func evolve(x int, y int, matIn [][]int, matOut [][]int, wg *sync.WaitGroup) {
	// Start with an active count of 0
	defer wg.Done()

	activeCells := 0
	yMax := len(matIn)
	xMax := len(matIn[0])

	// Need to iterate over 3x3 region centered at (x,y) and sum

	if y-1 >= 0 {
		if x-1 >= 0 {
			activeCells += matIn[y-1][x-1]
		}

		// UC for free (No need to check x bound here)
		activeCells += matIn[y-1][x]

		if x+1 < xMax {
			activeCells += matIn[y-1][x+1]
		}
	}

	if x-1 >= 0 {
		activeCells += matIn[y][x-1]
	}

	if x+1 < xMax {
		activeCells += matIn[y][x+1]
	}

	if y+1 < yMax {
		if x-1 >= 0 {
			activeCells += matIn[y+1][x-1]
		}

		// LC for free (No need to check x bound here)
		activeCells += matIn[y+1][x]

		if x+1 < xMax {
			activeCells += matIn[y+1][x+1]
		}
	}

	if activeCells > 3 || activeCells < 2 {
		matOut[y][x] = 0
	} else {
		if activeCells == 3 {
			matOut[y][x] = 1
		} else {
			matOut[y][x] = matIn[y][x]
		}
	}
}

// // MinMax is a one-pass min and max finder for slices
// func MinMax(arr []int) (int, int) {
// 	min := arr[0]
// 	max := arr[0]

// 	for _, val := range arr {
// 		if val < min {
// 			min = val
// 		}
// 		if val > max {
// 			max = val
// 		}
// 	}

// 	return min, max
// }
