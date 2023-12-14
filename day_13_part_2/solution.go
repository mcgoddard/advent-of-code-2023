package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

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
	puzzles := strings.Split(contentString, "\n\n")
	horizontalReflections := make(map[int]int)
	verticalReflections := make(map[int]int)
	for puzzleIndex, puzzle := range puzzles {
		puzzleLines := strings.Split(puzzle, "\n")
		// Check for a horizontal reflection
		matchIndex := findReflection(puzzleLines)
		if matchIndex > -1 {
			horizontalReflections[puzzleIndex] = matchIndex
			continue
		}
		// Invert puzzle and check for a vertical reflection
		puzzleColumns := make([][]rune, len(puzzleLines[0]))
		for i := range puzzleLines {
			for j := range puzzleLines[i] {
				if i == 0 {
					puzzleColumns[j] = make([]rune, len(puzzleLines))
				}
				puzzleColumns[j][i] = []rune(puzzleLines[i])[j]
			}
		}
		puzzleColumnStrings := make([]string, len(puzzleColumns))
		for i, column := range puzzleColumns {
			puzzleColumnStrings[i] = string(column)
		}
		matchIndex = findReflection(puzzleColumnStrings)
		if matchIndex > -1 {
			verticalReflections[puzzleIndex] = matchIndex
		} else {
			fmt.Println("No reflection found for puzzle!", puzzleIndex)
			panic("No reflection")
		}
	}
	// Print the result
	output := 0
	for i := range puzzles {
		if value, exists := horizontalReflections[i]; exists {
			output += (100 * (value + 1))
		} else if value, exists := verticalReflections[i]; exists {
			output += value + 1
		}
	}
	fmt.Println(output)
}

func findReflection(puzzleLines []string) int {
	lineCounts := make([]uint32, 0)
	for _, puzzleLine := range puzzleLines {
		lineCounts = append(lineCounts, convertToFlags(puzzleLine))
	}
	matchIndex := -1
	for i := 0; i < len(lineCounts)-1; i++ {
		diff := difference(lineCounts[i], lineCounts[i+1])
		if oneMistake := isPowerOfTwo(diff); diff == 0 || oneMistake {
			mistakes := 0
			if oneMistake {
				mistakes++
			}
			for j := 1; i-j >= 0 && i+j+1 < len(lineCounts); j++ {
				diff := difference(lineCounts[i-j], lineCounts[i+j+1])
				if diff != 0 {
					mistakes++
				}
			}
			if mistakes == 1 {
				matchIndex = i
			}
		}
	}
	return matchIndex
}

func convertToFlags(input string) uint32 {
	lineCount := uint32(0)
	for i, char := range input {
		if char == '.' {
			lineCount += uint32(math.Pow(2, float64(i)))
		}
	}
	return lineCount
}

func isPowerOfTwo(input uint32) bool {
	return input != 0 && (input&(input-1)) == 0
}

func difference(i1 uint32, i2 uint32) uint32 {
	return i1 ^ i2
}
