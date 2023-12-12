package main

import (
	"flag"
	"fmt"
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

const expansion = -1
const expansionValue = 1000000

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
		for x, char := range line {
			if char == '.' {
				newLine = append(newLine, 0)
			} else {
				newLine = append(newLine, galaxyNum)
				galaxyNum++
				galaxyInRow = true
				galaxyInColumn[x] = true
			}
		}
		if !galaxyInRow {
			for x := range line {
				newLine[x] = expansion
			}
		}
		galaxy = append(galaxy, newLine)
	}
	// Expand galaxy columns
	newGalaxy := make([][]int, 0)
	for i := range galaxy {
		newGalaxyLine := make([]int, 0)
		for j := range galaxy[i] {
			if galaxyInColumn[j] {
				newGalaxyLine = append(newGalaxyLine, galaxy[i][j])
			} else {
				newGalaxyLine = append(newGalaxyLine, expansion)
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
		output += distance(newGalaxy, galaxyPositions[pair.g1], galaxyPositions[pair.g2])
	}
	// Print the result
	fmt.Println(output)
}

func distance(galaxy [][]int, p1 Point, p2 Point) int {
	xDiff := 0
	for i := min(p1.x, p2.x); i < max(p1.x, p2.x); i++ {
		gridValue := galaxy[p1.y][i]
		if gridValue == expansion {
			xDiff += expansionValue
		} else {
			xDiff += 1
		}
	}
	yDiff := 0
	for i := min(p1.y, p2.y); i < max(p1.y, p2.y); i++ {
		gridValue := galaxy[i][p1.x]
		if gridValue == expansion {
			yDiff += expansionValue
		} else {
			yDiff += 1
		}
	}
	return xDiff + yDiff
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
