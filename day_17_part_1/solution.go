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
	x                int
	y                int
	vector           Vector
	alreadyTravelled int
}

type Vector struct {
	x int
	y int
}

type Node struct {
	f int
	g int
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
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len([]rune(line)))
		for j, char := range line {
			heatLoss, _ := strconv.Atoi(string(char))
			grid[i][j] = heatLoss
		}
	}
	// Calculate route
	open := map[Point]Node{
		{
			x:                0,
			y:                0,
			vector:           Vector{x: 1, y: 0},
			alreadyTravelled: 0,
		}: {
			f: 0,
			g: 0,
		},
		{
			x:                0,
			y:                0,
			vector:           Vector{x: 0, y: 1},
			alreadyTravelled: 0,
		}: {
			f: 0,
			g: 0,
		},
	}
	endPoint := Point{
		x: len(grid[0]) - 1,
		y: len(grid) - 1,
	}
	closed := make(map[Point]Node)
	output := 0
	parents := make(map[Point]Point)
	// var finalPoint Point
	for len(open) > 0 {
		// Find the lowest scoring point
		minScore := math.MaxInt
		var currentPoint Point
		var currentNode Node
		for p, n := range open {
			if n.f < minScore {
				minScore = n.f
				currentPoint = p
				currentNode = n
			}
		}
		// Remove from open
		delete(open, currentPoint)
		// Place on closed list
		closed[currentPoint] = currentNode
		// If we've just added the target to the closed list exit
		if currentPoint.x == endPoint.x && currentPoint.y == endPoint.y {
			output = currentNode.f
			// finalPoint = currentPoint
			break
		}
		// Get children
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				// Skip if not one of the valid directions
				if i == j || i*-1 == j {
					continue
				}
				newVector := Vector{x: j, y: i}
				newTravelled := 0
				if newVector == currentPoint.vector {
					newTravelled = currentPoint.alreadyTravelled + 1
				}
				childPoint := Point{
					x:                currentPoint.x + j,
					y:                currentPoint.y + i,
					vector:           newVector,
					alreadyTravelled: newTravelled,
				}
				// Skip if off the grid
				if childPoint.x < 0 || childPoint.x >= len(grid[0]) || childPoint.y < 0 || childPoint.y >= len(grid) {
					continue
				}
				// Skip if already evaluated
				if _, exists := closed[childPoint]; exists {
					continue
				}
				// Skip if 180
				if (newVector.x != 0 && newVector.x*-1 == currentPoint.vector.x) || (newVector.y != 0 && newVector.y*-1 == currentPoint.vector.y) {
					continue
				}
				// Skip if same direction and more than 3 steps
				if newVector == currentPoint.vector && currentPoint.alreadyTravelled >= 2 {
					continue
				}
				// A*
				g := currentNode.g + grid[childPoint.y][childPoint.x]
				h := int(math.Abs(float64(endPoint.x)-float64(childPoint.x)) + math.Abs(float64(endPoint.y)-float64(childPoint.y)))
				f := g + h
				// Skip if we hit a space not via a shorter route
				if previousNode, exists := open[childPoint]; exists && previousNode.g < g {
					continue
				}
				// Add the new node to the open list (or update if route shorter)
				open[childPoint] = Node{
					f: f,
					g: g,
				}
				parents[childPoint] = currentPoint
			}
		}
	}
	// Print the result
	fmt.Println(output)
}
