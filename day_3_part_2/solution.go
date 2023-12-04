package main

import (
	"flag"
	"fmt"
	"os"
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
	output := 0
	for i, line := range lines {
		for j, char := range line {
			char := string(char)
			if char == "*" {
				// Look for numbers around the *
				numbers := []int{}
				for scanLine := Max(0, i-1); scanLine <= Min(i+1, len(lines)-1); scanLine++ {
					for scanChar := 0; scanChar <= len(lines[i])-1; scanChar++ {
						startIndex := scanChar
						endIndex := scanChar + 1
						numberFound := false
						if _, err := strconv.Atoi(string(lines[scanLine][scanChar])); err == nil && string(lines[scanLine][scanChar]) != "." {
							numberFound = true
							for {
								if endIndex == len(lines[scanLine]) {
									break
								}
								if _, err := strconv.Atoi(string(lines[scanLine][endIndex])); err != nil || string(lines[scanLine][endIndex]) == "." {
									break
								}
								endIndex += 1
							}
						}
						if numberFound && (j-1 < endIndex && j+1 >= startIndex) {
							number, _ := strconv.Atoi(string(lines[scanLine][startIndex:endIndex]))
							scanChar = endIndex
							numbers = append(numbers, number)
						}
					}
				}
				if len(numbers) == 2 {
					gearRatio := numbers[0] * numbers[1]
					output += gearRatio
				}
			}
		}
	}
	// Print the result
	fmt.Println(output)
}
