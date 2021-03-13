package golife

import (
	"fmt"
	"sync"
)

func main() {
	const iterations int = 3

	currentState := [][]int{
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 1, 0, 0, 1},
	}

	fmt.Println(currentState)

	var wg sync.WaitGroup

	height := len(currentState)
	width := len(currentState[0])

	for i := 0; i < iterations; i++ {
		duplicate := make([][]int, len(currentState))
		for i := range currentState {
			duplicate[i] = make([]int, len(currentState[i]))
			copy(duplicate[i], currentState[i])
		}
		for j := 0; j < height; j++ {
			for k := 0; k < width; k++ {
				go evolve(k, j, currentState, duplicate, &wg)
			}
		}
		wg.Wait()

		currentState := duplicate
		duplicate = nil
	}

	fmt.Println(currentState)
}
