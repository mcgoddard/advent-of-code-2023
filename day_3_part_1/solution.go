package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
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
	content_string := string(content)
	// Split on newline
	lines := strings.Split(content_string, "\n")
	// For each number
	nonSymbols := []string{".", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	output := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			char := line[j]
			if _, err := strconv.Atoi(string(char)); err == nil {
				// Extract the complete number
				startIndex := j
				endIndex := j + 1
				for {
					if endIndex == len(lines[i]) {
						break
					}
					if _, err := strconv.Atoi(string(lines[i][endIndex])); err != nil {
						break
					}
					endIndex += 1
				}
				number, _ := strconv.Atoi(string(lines[i][startIndex:endIndex]))
				// Ensure we continue the scan after the current number
				j = endIndex
				// Check if there are any symbols around it
				symbolFound := false
				for scanLine := Max(0, i-1); scanLine <= Min(i+1, len(lines)-1); scanLine++ {
					if symbolFound {
						break
					}
					for scanChar := Max(0, startIndex-1); scanChar <= Min(endIndex, len(lines[i])-1); scanChar++ {
						if !slices.Contains(nonSymbols, string(lines[scanLine][scanChar])) {
							symbolFound = true
							break
						}
					}
				}
				if symbolFound {
					fmt.Println("Part number found: ", number)
					output += number
				}
			}
		}
	}
	// Print the result
	fmt.Println(output)
}
