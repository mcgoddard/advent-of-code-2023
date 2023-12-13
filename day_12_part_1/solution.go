package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Condition rune

const (
	workingChar Condition = '.'
	brokenChar  Condition = '#'
	unknownChar Condition = '?'
)

const (
	working = iota
	unknown = iota
	broken  = iota
)

type Entry struct {
	conditions []int
	counts     []int
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
	// Parse input
	lines := strings.Split(content_string, "\n")
	entries := make([]Entry, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		counts := strings.Split(parts[1], ",")
		countNums := make([]int, len(counts))
		for j := range counts {
			countNums[j], _ = strconv.Atoi(counts[j])
		}
		conditions := make([]int, len(parts[0]))
		for j, part := range parts[0] {
			switch Condition(part) {
			case workingChar:
				conditions[j] = working
			case brokenChar:
				conditions[j] = broken
			case unknownChar:
				conditions[j] = unknown
			default:
				fmt.Println("Unrecognised character", parts[0][j])
			}
		}
		entries[i] = Entry{
			conditions: conditions,
			counts:     countNums,
		}
	}
	// Calculate counts
	output := 0
	for _, entry := range entries {
		entryCount := count(entry.conditions, entry.counts)
		output += entryCount
	}
	// Print the result
	fmt.Println(output)
}

func count(conditions []int, counts []int) int {
	total := 0
	for _, block := range counts {
		total += block
	}
	brokenCount := 0
	notWorkingCount := 0
	for _, condition := range conditions {
		if condition != working {
			notWorkingCount++
			if condition == broken {
				brokenCount++
			}
		}
	}
	if brokenCount > total || notWorkingCount < total {
		return 0
	}
	if total == 0 {
		return 1
	}
	if conditions[0] == working {
		return count(conditions[1:], counts)
	}
	if conditions[0] == broken {
		length := counts[0]
		if matchBeginning(conditions, length) {
			if length == len(conditions) {
				return 1
			}
			return count(conditions[length+1:], counts[1:])
		}
		return 0
	}
	newConditions := append([]int{broken}, conditions[1:]...)
	return count(conditions[1:], counts) + count(newConditions, counts)
}

func matchBeginning(conditions []int, length int) bool {
	if len(conditions) != length && conditions[length] == broken {
		return false
	}
	for _, condition := range conditions[:length] {
		if condition == working {
			return false
		}
	}
	return true
}
