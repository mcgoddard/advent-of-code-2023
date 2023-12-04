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
	// Split on newline
	lines := strings.Split(content_string, "\n")
	// Get games
	output := 0
	for _, line := range lines {
		headerSplit := strings.Split(line, ": ")
		reveals := strings.Split(headerSplit[1], "; ")
		maxValues := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}
		for _, reveal := range reveals {
			colors := strings.Split(reveal, ", ")
			for _, color := range colors {
				parts := strings.Split(color, " ")
				if numberCubes, _ := strconv.Atoi(parts[0]); numberCubes > maxValues[parts[1]] {
					maxValues[parts[1]] = numberCubes
				}
			}
		}
		power := 1
		for _, count := range maxValues {
			power *= count
		}
		output += power
	}
	// Print the result
	fmt.Println(output)
}
