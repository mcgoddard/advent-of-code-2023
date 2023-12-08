package main

import (
	"flag"
	"fmt"
	"os"
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
	// Split on newline
	lines := strings.Split(content_string, "\n")
	// Read instructions list
	directionsMap := map[string]int{
		"L": 0,
		"R": 1,
	}
	instructions := strings.Split(lines[0], "")
	// Parse nodes
	nodes := make(map[string][]string, 0)
	for i := 2; i < len(lines); i++ {
		parts := strings.Split(lines[i], " = ")
		stripped := strings.ReplaceAll(parts[1], "(", "")
		stripped = strings.ReplaceAll(stripped, ")", "")
		nextNodes := strings.Split(stripped, ", ")
		nodes[parts[0]] = nextNodes
	}
	// Find start positions
	startPositions := make(map[string]void)
	for node := range nodes {
		if node[len(node)-1] == 'A' {
			startPositions[node] = member
		}
	}
	// Follow instructions
	stepsForStarts := []int{}
	for startPosition := range startPositions {
		currentNode := startPosition
		steps := 0
		for currentNode[len(currentNode)-1] != 'Z' {
			currentNode = nodes[currentNode][directionsMap[instructions[steps%len(instructions)]]]
			steps++
		}
		stepsForStarts = append(stepsForStarts, steps)
	}
	// Calculate Lowest Common Multiple to find number of steps where all routes landed on Z
	output := (stepsForStarts[0] * stepsForStarts[1]) / greatestCommonDivisor(stepsForStarts[0], stepsForStarts[1])
	for i := 2; i < len(stepsForStarts); i++ {
		output = (output * stepsForStarts[i]) / greatestCommonDivisor(output, stepsForStarts[i])
	}
	// Print the result
	fmt.Println(output)
}

func greatestCommonDivisor(a, b int) int {
	for b > 0 {
		c := b
		b = a % b
		a = c
	}
	return a
}
