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
	const iterations int = 1500

	currentState := [][]int{}
	height := 10 + rand.Intn(30)
	width := 10 + rand.Intn(50)

	for i := 0; i < height; i++ {
		row := []int{}
		for j := 0; j < width; j++ {
			if rand.Float32() < 0.45 {
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

	totalTime := 0
	totalIterations := 0
	startTime := time.Now()

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

		totalTime += int(time.Now().Sub(startTime).Milliseconds())
		startTime = time.Now()
		totalIterations++

		if MatsEqual(currentState, duplicate) {
			fmt.Println()
			fmt.Println("Have reached a terminal state...")
			fmt.Println()
			break
		}
		currentState = duplicate
		duplicate = nil
		PrintMatrix(currentState)
		fmt.Println()
		fmt.Println("Current Evolution: ", iter)
		fmt.Println("Number of active workers: ", runtime.NumGoroutine())
		time.Sleep(50 * time.Millisecond)
		CallClear()
	}
	wg.Wait()
	fmt.Println("Ran", totalIterations, "iterations in", totalTime, "milliseconds (CPU time)...")
}

// MatsEqual takes in two 2D slices and compares for equality
func MatsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for aIndex, row := range a {
		if len(row) != len(b[aIndex]) {
			return false
		}
		for rowIndex, val := range row {
			if val != b[aIndex][rowIndex] {
				return false
			}
		}
	}
	return true
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
