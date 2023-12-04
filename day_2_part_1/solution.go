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
	maxValues := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	output := 0
	for _, line := range lines {
		headerSplit := strings.Split(line, ": ")
		gameId, _ := strconv.Atoi(strings.Replace(headerSplit[0], "Game ", "", 1))
		reveals := strings.Split(headerSplit[1], "; ")
		gameValid := true
		for _, reveal := range reveals {
			if !gameValid {
				break
			}
			colors := strings.Split(reveal, ", ")
			for _, color := range colors {
				parts := strings.Split(color, " ")
				if numberCubes, _ := strconv.Atoi(parts[0]); numberCubes > maxValues[parts[1]] {
					gameValid = false
					break
				}
			}
		}
		if gameValid {
			output += gameId
		}
	}
	// Print the result
	fmt.Println(output)
}
