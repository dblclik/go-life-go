package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func main() {
	const iterations int = 50

	currentState := [][]int{}
	height := 10 + rand.Intn(30)
	width := 10 + rand.Intn(50)

	for i := 0; i < height; i++ {
		row := []int{}
		for j := 0; j < width; j++ {
			if rand.Float32() < 0.5 {
				row = append(row, 0)
			} else {
				row = append(row, 1)
			}
		}
		currentState = append(currentState, row)
	}

	PrintMatrix(currentState)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)
	CallClear()

	var wg sync.WaitGroup

	for iter := 0; iter < iterations; iter++ {
		duplicate := make([][]int, len(currentState))
		for i := range currentState {
			duplicate[i] = make([]int, len(currentState[i]))
			copy(duplicate[i], currentState[i])
		}
		for j := 0; j < height; j++ {
			for k := 0; k < width; k++ {
				wg.Add(1)
				go Evolve(k, j, currentState, duplicate, &wg)
			}
		}

		currentState = duplicate
		duplicate = nil
		PrintMatrix(currentState)
		fmt.Println()
		time.Sleep(50 * time.Millisecond)
		CallClear()
	}
	wg.Wait()
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// CallClear lets us clear the terminal for easier printing of the matrices
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// PrintMatrix pretty prints the given matrix
func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// Evolve iterates the matrix forward by 1 step
func Evolve(x int, y int, matIn [][]int, matOut [][]int, wg *sync.WaitGroup) {
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
			matOut[y][x] = 1 * matIn[y][x]
		}
	}
}
