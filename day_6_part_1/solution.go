package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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
	// Split on newline
	type Race struct {
		time     int
		distance int
	}
	lines := strings.Split(content_string, "\n")
	output := 1
	times := strings.Split(lines[0], " ")
	filteredTimes := make([]int, 0)
	for _, time := range times {
		if timeInt, err := strconv.Atoi(time); err == nil {
			filteredTimes = append(filteredTimes, timeInt)
		}
	}
	distances := strings.Split(lines[1], " ")
	filteredDistances := make([]int, 0)
	for _, distance := range distances {
		if distanceInt, err := strconv.Atoi(distance); err == nil {
			filteredDistances = append(filteredDistances, distanceInt)
		}
	}
	fmt.Println("Times", filteredTimes)
	fmt.Println("Distances", filteredDistances)
	races := make([]Race, len(filteredTimes))
	for i := range filteredTimes {
		races[i] = Race{time: filteredTimes[i], distance: filteredDistances[i]}
	}
	fmt.Println("Races", races)
	// Brute force each option for winning in each race
	for _, race := range races {
		winCount := 0
		for buttonHeld := 0; buttonHeld < race.time; buttonHeld++ {
			if buttonHeld*(race.time-buttonHeld) > race.distance {
				winCount++
			}
		}
		output *= winCount
	}
	// Print the result
	fmt.Println(output)
}
