package main

import (
	"flag"
	"fmt"
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
	cardCounts := make(map[int]int)
	lines := strings.Split(content_string, "\n")
	output := 0
	// Initialise counts at 1 for all cards
	for i := range lines {
		cardCounts[i+1] = 1
	}
	// Read the lines
	for _, line := range lines {
		ticketCount := 0
		lineHeader := strings.Split(line, ":")
		header := strings.Split(lineHeader[0], " ")
		cardId, _ := strconv.Atoi(header[len(header)-1])
		numbersSplit := strings.Split(lineHeader[1], "|")
		// Build a set of winning numbers
		winners := make(map[int]void)
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
				if _, exists := winners[numberInt]; exists {
					ticketCount += 1
				}
			}
		}
		// Bump each of the winning tickets (as many times as we have of this ticket)
		if ticketCount > 0 {
			for j := 1; j <= ticketCount && cardId+j <= len(lines); j++ {
				cardCounts[cardId+j] += cardCounts[cardId]
			}
		}
		// Add our current card count to the total
		output += cardCounts[cardId]
	}
	// Print the result
	fmt.Println(output)
}
