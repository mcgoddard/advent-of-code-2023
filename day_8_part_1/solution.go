package main

import (
	"flag"
	"fmt"
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
	content_string := string(content)
	output := 0
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
	// Follow instructions
	position := "AAA"
	finalDestination := "ZZZ"
	for {
		position = nodes[position][directionsMap[instructions[output%len(instructions)]]]
		output++
		if position == finalDestination {
			break
		}
	}
	// Print the result
	fmt.Println(output)
}
