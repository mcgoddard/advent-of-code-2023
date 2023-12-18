package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
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
	lines := strings.Split(contentString, "\n")
	verticies := []Point{{x: 0, y: 0}}
	for i, line := range lines {
		parts := strings.Split(line, " ")
		meters64, _ := strconv.ParseInt(string([]rune(parts[2])[2:7]), 16, 32)
		meters := int(meters64)
		directionNum, _ := strconv.Atoi(string([]rune(parts[2])[7]))
		previous := verticies[i]
		modifyX := 0
		modifyY := 0
		if directionNum == 2 {
			modifyX = -1 * meters
		} else if directionNum == 0 {
			modifyX = meters
		} else if directionNum == 3 {
			modifyY = -1 * meters
		} else if directionNum == 1 {
			modifyY = meters
		}
		newVertex := Point{
			x: previous.x + modifyX,
			y: previous.y + modifyY,
		}
		verticies = append(verticies, newVertex)
	}
	// Shoelace for area
	output := Shoelace(verticies)
	// Add half the perimeter distance
	perimeter := 0
	for i, vertex := range verticies[:len(verticies)-1] {
		perimeter += Distance(vertex, verticies[i+1])
	}
	perimeter += Distance(verticies[len(verticies)-1], verticies[0])
	output += perimeter/2 + 1
	// Print the result
	fmt.Println(output)
}

func Shoelace(verticies []Point) int {
	numberOfVerticies := len(verticies)
	twiceArea := 0
	for i, vertex := range verticies[:numberOfVerticies-1] {
		twiceArea += (vertex.x * verticies[i+1].y) - (vertex.y * verticies[i+1].x)
	}
	twiceArea += (verticies[numberOfVerticies-1].x * verticies[0].y) - (verticies[numberOfVerticies-1].y * verticies[0].x)
	return int(math.Abs(float64(twiceArea))) / 2
}

func Distance(a Point, b Point) int {
	xDiff := int(math.Abs(float64(b.x - a.x)))
	yDiff := int(math.Abs(float64(b.y - a.y)))
	return xDiff + yDiff
}
