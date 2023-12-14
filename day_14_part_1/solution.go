package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type TileType int

const (
	empty  TileType = iota
	square TileType = iota
	round  TileType = iota
)

var rockMap = map[rune]TileType{
	'O': round,
	'#': square,
	'.': empty,
}

func main() {
	exampleFlagPtr := flag.Bool("example", false, "Use the example file if set")
	flag.Parse()
	// Read input file
	inputFile := "input.txt"
	if *exampleFlagPtr {
		inputFile = "example.txt"
	}
	content, _ := os.ReadFile(inputFile)
	contentString := string(content)
	// Parse input
	lines := strings.Split(contentString, "\n")
	// Rotate grid
	grid := rotateGrid(lines)
	// Calculate load
	output := 0
	for _, row := range grid {
		output += calculateRow(row)
	}
	// Print the result
	fmt.Println(output)
}

func rotateGrid(lines []string) [][]TileType {
	grid := make([][]TileType, len(lines[0]))
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if i == 0 {
				grid[j] = make([]TileType, len(lines))
			}
			grid[j][i] = rockMap[[]rune(lines[i])[j]]
		}
	}
	return grid
}

func calculateRow(row []TileType) int {
	firstSquareIndex := -1
	roundCount := 0
	for i, value := range row {
		if value == square {
			firstSquareIndex = i
			break
		} else if value == round {
			roundCount++
		}
	}
	result := 0
	for i := 0; i < roundCount; i++ {
		result += len(row) - i
	}
	if firstSquareIndex == -1 {
		return result
	}
	return result + calculateRow(row[firstSquareIndex+1:])
}
