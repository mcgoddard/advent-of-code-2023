package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

type Point struct {
	value  int
	start  bool
	mapped bool
}

type By func(p1, p2 *Point) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(points []Point) {
	ps := &pointSorter{
		points: points,
		by:     by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type pointSorter struct {
	points []Point
	by     func(p1, p2 *Point) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *pointSorter) Len() int {
	return len(s.points)
}

// Swap is part of sort.Interface.
func (s *pointSorter) Swap(i, j int) {
	s.points[i], s.points[j] = s.points[j], s.points[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *pointSorter) Less(i, j int) bool {
	return s.by(&s.points[i], &s.points[j])
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
			end:   start + seedRange - 1,
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
				{start: sourceRangeStart, end: sourceRangeStart + rangeLength - 1},
				{start: destinationRangeStart, end: destinationRangeStart + rangeLength - 1},
			})
			lineCount++
		}
		i += lineCount
	}
	fmt.Println(seeds)
	fmt.Println(lists)
	// Check each seed
	output := int((^uint(0)) >> 1)
	ranges := seeds
	for _, list := range lists {
		newRanges := splitRanges(list, ranges)
		ranges = newRanges
	}
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
	for _, valueRange := range ranges {
		points := make([]Point, 0)
		points = append(points, Point{
			value:  valueRange.start,
			start:  true,
			mapped: false,
		})
		points = append(points, Point{
			value:  valueRange.end,
			start:  false,
			mapped: false,
		})
		for _, mapRange := range list {
			points = append(points, Point{
				value:  mapRange[0].start,
				start:  true,
				mapped: true,
			})
			points = append(points, Point{
				value:  mapRange[0].end,
				start:  false,
				mapped: true,
			})
		}
		value := func(p1, p2 *Point) bool {
			if p1.value == p2.value {
				if p1.start && !p2.start {
					return false
				} else if !p1.start && p2.start {
					return true
				}
				return p1.mapped
			}
			return p1.value < p2.value
		}
		By(value).Sort(points)
		// fmt.Println("Value range: ", valueRange)
		start := 0
		inRange := false
		inMapped := false
		lastEnd := -1
		for _, point := range points {
			// fmt.Print(point)
			if !inRange {
				if point.start {
					if !point.mapped {
						// fmt.Print(" open range")
						start = point.value
						inRange = true
					} else {
						// fmt.Print(" opening mapped")
						inMapped = true
					}
				} else {
					// fmt.Print(" closing mapped")
					inMapped = false
				}
			} else {
				// We're in a range, is this the end of it?
				if !point.start {
					// Is this a map end or a value end?
					if !point.mapped {
						// Value end
						// fmt.Print(" closing range")
						inRange = false
						if !inMapped {
							if point.value > lastEnd {
								// fmt.Print(" closing range unmapped and starting mapped range")
								newRanges = append(newRanges, Range{
									start: start,
									end:   point.value,
								})
								lastEnd = point.value
							}
						} else {
							for i, mapRange := range list {
								if start >= mapRange[0].start && start <= mapRange[0].end {
									if point.value > lastEnd {
										// fmt.Print(" closing range mapped and starting mapped range ", mapRange[1].start, mapRange[0].start, start)
										newRanges = append(newRanges, Range{
											start: mapRange[1].start - mapRange[0].start + start,
											end:   mapRange[1].start - mapRange[0].start + point.value,
										})
										lastEnd = point.value
										break
									}
								}
								if i == len(list)-1 {
									// fmt.Print(" failed to find range")
								}
							}
						}
					} else {
						// Map end
						inMapped = false
						for i, mapRange := range list {
							if start >= mapRange[0].start && start <= mapRange[0].end {
								if point.value > lastEnd {
									// fmt.Print(" closing range mapped and starting unmapped range ", mapRange[1].start, mapRange[0].start, start)
									newRanges = append(newRanges, Range{
										start: mapRange[1].start - mapRange[0].start + start,
										end:   mapRange[1].start - mapRange[0].start + point.value,
									})
									lastEnd = point.value
									break
								}
							}
							if i == len(list)-1 {
								// fmt.Print(" failed to find range")
							}
						}
						start = point.value + 1
					}
				} else { // We're in a range and have hit a start, switch to map
					if point.value > lastEnd {
						// fmt.Print(" closing range unmapped open mapped range")
						newRanges = append(newRanges, Range{
							start: start,
							end:   point.value - 1,
						})
						lastEnd = point.value
						inMapped = true
						start = point.value
					}
				}
			}
			// fmt.Println()
		}
	}
	return newRanges
}
