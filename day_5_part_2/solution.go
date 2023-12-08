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
	seeds := make([]Range, len(seedStrings)/2)
	for i := 0; i < len(seedStrings); i += 2 {
		start, _ := strconv.Atoi(seedStrings[i])
		seedRange, _ := strconv.Atoi(seedStrings[i+1])
		seeds[i/2] = Range{
			start: start,
			end:   start + seedRange,
		}
	}
	// Extract lists
	lists := make([][][]Range, 0)
	for i := 2; i < len(lines); i++ {
		lineCount := 1
		lists = append(lists, make([][]Range, 0))
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
			lists[len(lists)-1] = append(lists[len(lists)-1], []Range{
				{start: sourceRangeStart, end: sourceRangeStart + rangeLength},
				{start: destinationRangeStart, end: destinationRangeStart + rangeLength},
			})
			lineCount++
		}
		i += lineCount
	}
	// Check each seed
	output := int((^uint(0)) >> 1)
	ranges := seeds
	for _, list := range lists {
		ranges = splitRanges(list, ranges)
	}
	// Get the min
	for _, valueRanges := range ranges {
		if valueRanges.start < output {
			output = valueRanges.start
		}
	}
	// Print the result
	fmt.Println(output)
}

func splitRanges(list [][]Range, ranges []Range) []Range {
	newRanges := make([]Range, 0)
	for _, currentRange := range list {
		sourceRange := currentRange[0]
		destinationRange := currentRange[1]
		tempIntervals := make([]Range, 0)

		for _, sourceValues := range ranges {
			left := []int{sourceValues.start, min(sourceValues.end, sourceRange.start)}
			mid := []int{max(sourceValues.start, sourceRange.start), min(sourceValues.end, sourceRange.end)}
			right := []int{max(sourceRange.end, sourceValues.start), sourceValues.end}
			if left[1] > left[0] {
				tempIntervals = append(tempIntervals, Range{start: left[0], end: left[1]})
			}
			if mid[1] > mid[0] {
				newRanges = append(
					newRanges,
					Range{start: mid[0] - sourceRange.start + destinationRange.start, end: mid[1] - sourceRange.start + destinationRange.start},
				)
			}
			if right[1] > right[0] {
				tempIntervals = append(tempIntervals, Range{start: right[0], end: right[1]})
			}
		}
		ranges = tempIntervals
	}
	return append(newRanges, ranges...)
}
