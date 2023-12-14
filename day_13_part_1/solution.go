package main

import (
	"flag"
	"fmt"
	"hash/fnv"
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
	lineHashes := make([]uint32, 0)
	for _, puzzleLine := range puzzleLines {
		h := fnv.New32a()
		h.Write([]byte(puzzleLine))
		lineHashes = append(lineHashes, h.Sum32())
	}
	matchIndex := -1
	for i := 0; i < len(lineHashes)-1; i++ {
		if lineHashes[i] == lineHashes[i+1] {
			matching := true
			for j := 1; i-j >= 0 && i+j+1 < len(lineHashes); j++ {
				if lineHashes[i-j] != lineHashes[i+j+1] {
					matching = false
				}
			}
			if matching {
				matchIndex = i
			}
		}
	}
	return matchIndex
}
