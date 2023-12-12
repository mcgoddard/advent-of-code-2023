package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pair struct {
	g1 int
	g2 int
}

type Point struct {
	x int
	y int
}

func main() {
	exampleFlagPtr := flag.Bool("example", false, "Use the example file if set")
	example2FlagPtr := flag.Bool("example2", false, "Use the second example file if set")
	flag.Parse()
	// Read input file
	inputFile := "input.txt"
	if *exampleFlagPtr {
		inputFile = "example.txt"
	} else if *example2FlagPtr {
		inputFile = "example2.txt"
	}
	content, _ := os.ReadFile(inputFile)
	content_string := string(content)
	// Split on newline
	lines := strings.Split(content_string, "\n")
	galaxyNum := 1
	galaxy := make([][]int, 0)
	galaxyInColumn := make([]bool, len(lines[0]))
	for _, line := range lines {
		newLine := make([]int, 0)
		galaxyInRow := false
		for y, char := range line {
			if char == '.' {
				newLine = append(newLine, 0)
			} else {
				newLine = append(newLine, galaxyNum)
				galaxyNum++
				galaxyInRow = true
				galaxyInColumn[y] = true
			}
		}
		galaxy = append(galaxy, newLine)
		if !galaxyInRow {
			anotherNewLine := make([]int, len(newLine))
			copy(anotherNewLine, newLine)
			galaxy = append(galaxy, anotherNewLine)
		}
	}
	// Expand galaxy columns
	newGalaxy := make([][]int, 0)
	for i := range galaxy {
		newGalaxyLine := make([]int, 0)
		for j := range galaxy[i] {
			newGalaxyLine = append(newGalaxyLine, galaxy[i][j])
			if !galaxyInColumn[j] {
				newGalaxyLine = append(newGalaxyLine, 0)
			}
		}
		newGalaxy = append(newGalaxy, newGalaxyLine)
	}
	// Set galaxy positions
	galaxyPositions := make(map[int]Point, 0)
	for y := range newGalaxy {
		for x := range newGalaxy[y] {
			if newGalaxy[y][x] != 0 {
				galaxyPositions[newGalaxy[y][x]] = Point{x: x, y: y}
			}
		}
	}
	// Get pairs
	pairs := make([]Pair, 0)
	for i := 1; i < galaxyNum; i++ {
		for j := i + 1; j < galaxyNum; j++ {
			pairs = append(pairs, Pair{g1: i, g2: j})
		}
	}
	// Calculate result
	output := 0
	for _, pair := range pairs {
		output += cartesianDistance(galaxyPositions[pair.g1], galaxyPositions[pair.g2])
	}
	// Print the result
	fmt.Println(output)
}

func cartesianDistance(p1 Point, p2 Point) int {
	xDiff := int(math.Abs(float64(p2.x - p1.x)))
	yDiff := int(math.Abs(float64(p2.y - p1.y)))
	return xDiff + yDiff
}
