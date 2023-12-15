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
	contentString := string(content)
	// Parse input
	steps := strings.Split(contentString, ",")
	output := 0
	for _, step := range steps {
		output += hashAlgorithm(step)
	}
	// Print the result
	fmt.Println(output)
}

func hashAlgorithm(step string) int {
	currentValue := 0
	for _, char := range step {
		currentValue += int(char)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}
