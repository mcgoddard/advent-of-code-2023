package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
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
	times := strings.Split(lines[0], " ")
	filteredTimes := make([]string, 0)
	for _, time := range times {
		if _, err := strconv.Atoi(time); err == nil {
			filteredTimes = append(filteredTimes, time)
		}
	}
	timeString := ""
	for _, time := range filteredTimes {
		timeString += time
	}
	finalTime, _ := strconv.Atoi(timeString)
	distances := strings.Split(lines[1], " ")
	filteredDistances := make([]string, 0)
	for _, distance := range distances {
		if _, err := strconv.Atoi(distance); err == nil {
			filteredDistances = append(filteredDistances, distance)
		}
	}
	distanceString := ""
	for _, distance := range filteredDistances {
		distanceString += distance
	}
	finalDistance, _ := strconv.Atoi(distanceString)
	fmt.Println("Time", finalTime)
	fmt.Println("Distance", finalDistance)
	race := Race{time: finalTime, distance: finalDistance}
	fmt.Println("Race", race)
	// Calculate output
	quadratic := int(math.Sqrt(math.Pow(float64(race.time), 2) - float64(4*race.distance)))
	low := ((-1 * race.time) - quadratic) / 2
	high := ((-1 * race.time) + quadratic) / 2
	output := high - low + 1
	// output := 0
	// for buttonHeld := 0; buttonHeld < race.time; buttonHeld++ {
	// 	if buttonHeld*(race.time-buttonHeld) > race.distance {
	// 		output++
	// 	}
	// }
	// Print the result
	fmt.Println(output)
	elapsed := time.Since(start)
	fmt.Println("Took ", elapsed)
}
