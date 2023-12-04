package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type void struct{}

var member void

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
	for _, line := range lines {
		ticketCount := 0
		lineHeader := strings.Split(line, ":")
		numbersSplit := strings.Split(lineHeader[1], "|")
		winners := make(map[int]void)
		// Build a set of winning numbers
		for _, winner := range strings.Split(numbersSplit[0], " ") {
			if winner != "" {
				winnerInt, _ := strconv.Atoi(winner)
				winners[winnerInt] = member
			}
		}
		// Check each of our numbers against the winners set
		for _, number := range strings.Split(numbersSplit[1], " ") {
			if number != "" {
				numberInt, _ := strconv.Atoi(number)
				_, exists := winners[numberInt]
				if exists {
					ticketCount += 1
				}
			}
		}
		ticketValue := 0
		if ticketCount == 1 {
			ticketValue = 1
		} else if ticketCount > 1 {
			ticketValue = int(math.Pow(2, float64(ticketCount-1)))
		}
		output += ticketValue
	}
	// Print the result
	fmt.Println(output)
}
