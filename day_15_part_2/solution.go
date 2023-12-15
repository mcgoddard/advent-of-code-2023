package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	Label       string
	FocalLength int
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
	contentString := string(content)
	// Parse input
	steps := strings.Split(contentString, ",")
	output := 0
	boxes := make(map[int][]Lens)
	for _, step := range steps {
		operationIndex := findIndexString(step, '-')
		if operationIndex < 0 {
			operationIndex = findIndexString(step, '=')
		}
		operation := step[operationIndex]
		label := string([]rune(step)[:operationIndex])
		boxNum := hashAlgorithm(label)
		if _, exists := boxes[boxNum]; !exists {
			boxes[boxNum] = make([]Lens, 0)
		}
		labelIndex := findLensIndex(boxes[boxNum], label)
		if operation == '-' {
			if labelIndex >= 0 {
				boxes[boxNum] = append(boxes[boxNum][:labelIndex], boxes[boxNum][labelIndex+1:]...)
			}
		} else {
			focalLengthString := string([]rune(step)[operationIndex+1])
			focalLength, _ := strconv.Atoi(focalLengthString)
			if labelIndex >= 0 {
				boxes[boxNum][labelIndex] = Lens{
					Label:       label,
					FocalLength: focalLength,
				}
			} else {
				boxes[boxNum] = append(boxes[boxNum], Lens{
					Label:       label,
					FocalLength: focalLength,
				})
			}
		}
	}
	for boxNum, contents := range boxes {
		for i, lens := range contents {
			output += (boxNum + 1) * (i + 1) * lens.FocalLength
		}
	}
	// Print the result
	fmt.Println(output)
}

func findIndexString(s string, c rune) int {
	for i, character := range s {
		if character == c {
			return i
		}
	}
	return -1
}

func findLensIndex(box []Lens, label string) int {
	for i, lens := range box {
		if lens.Label == label {
			return i
		}
	}
	return -1
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
