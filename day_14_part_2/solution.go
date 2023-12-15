package main

import (
	"encoding/json"
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

const cycles = 1000000000

type Board [][]TileType

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
	grid := convertGrid(lines)
	// Run cycles
	i := 1
	seen := make(map[string]int)
	scores := make([]int, 1)
	boardKey := ""
	for i < cycles {
		// Rotate 4 times
		for j := 0; j < 4; j++ {
			grid = rotateGrid(grid, true)
		}
		// Calculate load
		score := 0
		for i, row := range grid {
			score += calculateRow(row, i+1)
		}
		// Add to score list and check if we've seen it before
		scores = append(scores, score)
		boardKeyBytes, _ := json.Marshal(grid)
		boardKey = string(boardKeyBytes)
		if _, exists := seen[boardKey]; exists {
			break
		}
		seen[boardKey] = i
		i++
	}
	// Extract the final score
	startOfLoopIndex := seen[boardKey]
	loopLength := i - startOfLoopIndex
	output := scores[startOfLoopIndex+((cycles-startOfLoopIndex)%loopLength)]
	// Print the result
	fmt.Println(output)
}

func convertGrid(lines []string) [][]TileType {
	grid := make([][]TileType, len(lines))
	for i := 0; i < len(lines); i++ {
		grid[i] = make([]TileType, len(lines[0]))
		for j := 0; j < len(lines[0]); j++ {
			grid[i][j] = rockMap[[]rune(lines[i])[j]]
		}
	}
	grid = rotateGrid(grid, false)
	grid = rotateGrid(grid, false)
	return grid
}

func rotateGrid(lines [][]TileType, move bool) [][]TileType {
	transposedGrid := transposeGrid(lines)
	reversedRows := reverseRows(transposedGrid)
	if !move {
		return reversedRows
	}
	movedGrid := make([][]TileType, len(reversedRows))
	for i, row := range reversedRows {
		movedGrid[i] = moveRow(row)
	}
	return movedGrid
}

func transposeGrid(lines [][]TileType) [][]TileType {
	transposedGrid := make([][]TileType, len(lines[0]))
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if i == 0 {
				transposedGrid[j] = make([]TileType, len(lines))
			}
			transposedGrid[j][i] = lines[i][j]
		}
	}
	return transposedGrid
}

func reverseRows(lines [][]TileType) [][]TileType {
	grid := make([][]TileType, len(lines))
	// Reverse each row
	for i := 0; i < len(lines); i++ {
		grid[i] = make([]TileType, len(lines[0]))
		for j := len(lines[0]); j > 0; j-- {
			grid[i][j-1] = lines[i][len(lines[0])-j]
		}
	}
	return grid
}

func calculateRow(row []TileType, score int) int {
	total := 0
	for _, value := range row {
		if value == round {
			total += score
		}
	}
	return total
}

func moveRow(row []TileType) []TileType {
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
	resultLength := firstSquareIndex
	if firstSquareIndex == -1 {
		resultLength = len(row)
	}
	result := make([]TileType, resultLength)
	for i := 0; i < roundCount; i++ {
		result[i] = round
	}
	if firstSquareIndex == -1 {
		return result
	}
	return append(append(result, square), moveRow(row[firstSquareIndex+1:])...)
}
