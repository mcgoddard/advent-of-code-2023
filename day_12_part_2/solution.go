package main

import (
	"encoding/json"
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
	Conditions []int
	Counts     []int
}

const copies = 5

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
		countNums := make([]int, len(counts)*copies)
		for k := 0; k < copies; k++ {
			for j := range counts {
				countNums[(k*len(counts))+j], _ = strconv.Atoi(counts[j])
			}
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
		for i := 0; i < copies-1; i++ {
			conditions = append(conditions, unknown)
			for j, part := range parts[0] {
				switch Condition(part) {
				case workingChar:
					conditions = append(conditions, working)
				case brokenChar:
					conditions = append(conditions, broken)
				case unknownChar:
					conditions = append(conditions, unknown)
				default:
					fmt.Println("Unrecognised character", parts[0][j])
				}
			}
		}
		entries[i] = Entry{
			Conditions: conditions,
			Counts:     countNums,
		}
	}
	// Calculate counts
	output := 0
	countCached = memoize(count)
	for _, entry := range entries {
		entryCount := countCached.call(entry.Conditions, entry.Counts)
		output += entryCount
	}
	// Print the result
	fmt.Println(output)
}

var countCached *memoized

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
		return countCached.call(conditions[1:], counts)
	}
	if conditions[0] == broken {
		length := counts[0]
		if matchBeginning(conditions, length) {
			if length == len(conditions) {
				return 1
			}
			return countCached.call(conditions[length+1:], counts[1:])
		}
		return 0
	}
	newConditions := append([]int{broken}, conditions[1:]...)
	return countCached.call(conditions[1:], counts) + countCached.call(newConditions, counts)
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

type memoized struct {
	f     func([]int, []int) int
	cache map[string]int
}

func memoize(f func([]int, []int) int) *memoized {
	return &memoized{f: f, cache: make(map[string]int)}
}

func (m *memoized) call(conditions []int, counts []int) int {
	entry := Entry{
		Conditions: conditions,
		Counts:     counts,
	}
	keyBytes, _ := json.Marshal(entry)
	key := string(keyBytes)
	if v, ok := m.cache[key]; ok {
		return v
	}
	result := m.f(conditions, counts)
	m.cache[key] = result
	return result
}
