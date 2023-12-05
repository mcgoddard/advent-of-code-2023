package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
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
	// Get seed numbers
	seedStrings := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seeds := make([]int, len(seedStrings))
	for i, seedString := range seedStrings {
		if seed, err := strconv.Atoi(seedString); err == nil {
			seeds[i] = seed
		}
	}
	// Extract maps
	maps := make(map[string]map[string]map[Range]int, 0)
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		mapTitle := strings.Split(line, " map:")[0]
		titleParts := strings.Split(mapTitle, "-")
		fromName := titleParts[0]
		toName := titleParts[2]
		lineCount := 1
		maps[fromName] = make(map[string]map[Range]int)
		maps[fromName][toName] = make(map[Range]int)
		for {
			if i+lineCount >= len(lines) {
				break
			}
			dataLine := lines[i+lineCount]
			if dataLine == "" {
				break
			}
			dataParts := strings.Split(dataLine, " ")
			destinationRangeStart, _ := strconv.Atoi(dataParts[0])
			sourceRangeStart, _ := strconv.Atoi(dataParts[1])
			rangeLength, _ := strconv.Atoi(dataParts[2])
			maps[fromName][toName][Range{start: sourceRangeStart, end: sourceRangeStart + rangeLength}] = destinationRangeStart
			lineCount++
		}
		i += lineCount
	}
	// Check each seed
	output := int((^uint(0)) >> 1)
	fmt.Println("Seeds ", seeds)
	for _, seed := range seeds {
		fromKey := "seed"
		currentValue := seed
		for {
			if _, exists := maps[fromKey]; !exists {
				break
			}
			for toKey, values := range maps[fromKey] {
				// If the value is mapped, take the new value, otherwise pass it through to the next map as is
				for rangeKey, startDestination := range values {
					if currentValue >= rangeKey.start && currentValue < rangeKey.end {
						currentValue = (currentValue - rangeKey.start) + startDestination
						break
					}
				}
				// Move to the next map
				fromKey = toKey
			}
		}
		if currentValue < output {
			output = currentValue
		}
	}
	// Print the result
	fmt.Println(output)
}
