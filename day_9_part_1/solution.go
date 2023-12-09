package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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
	content_string := string(content)
	output := 0
	// Split on newline
	lines := strings.Split(content_string, "\n")
	// Convert to numbers
	histories := make([][][]int, 0)
	for _, line := range lines {
		valueStrings := strings.Split(line, " ")
		newHistory := make([][]int, 1)
		for _, valueString := range valueStrings {
			value, _ := strconv.Atoi(valueString)
			newHistory[0] = append(newHistory[0], value)
		}
		histories = append(histories, newHistory)
	}
	// Process each history
	for k := range histories {
		i := 0
		reachedEnd := true
		for !reachedEnd || i == 0 {
			reachedEnd = true
			nextDataset := make([]int, len(histories[k][i])-1)
			for j := 0; j < len(nextDataset); j++ {
				nextDataset[j] = histories[k][i][j+1] - histories[k][i][j]
				if nextDataset[j] != 0 {
					reachedEnd = false
				}
			}
			histories[k] = append(histories[k], nextDataset)
			i++
		}
	}
	// Calculate next values
	for k := range histories {
		for i := len(histories[k]) - 1; i > 0; i-- {
			lowerLineFinal := histories[k][i][len(histories[k][i])-1]
			upperLineFinal := histories[k][i-1][len(histories[k][i-1])-1]
			histories[k][i-1] = append(histories[k][i-1], lowerLineFinal+upperLineFinal)
		}
		output += histories[k][0][len(histories[k][0])-1]
	}
	// Print the result
	fmt.Println(output)
}
