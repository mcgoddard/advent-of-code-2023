package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Tile string

const (
	dug   Tile = "#"
	undug Tile = "."
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
	leastX := math.MaxInt
	mostX := math.MinInt
	leastY := math.MaxInt
	mostY := math.MinInt
	for i, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		meters, _ := strconv.Atoi(parts[1])
		previous := verticies[i]
		modifyX := 0
		modifyY := 0
		if direction == "L" {
			modifyX = -1 * meters
		} else if direction == "R" {
			modifyX = meters
		} else if direction == "U" {
			modifyY = -1 * meters
		} else if direction == "D" {
			modifyY = meters
		}
		newVertex := Point{
			x: previous.x + modifyX,
			y: previous.y + modifyY,
		}
		verticies = append(verticies, newVertex)
		if newVertex.x < leastX {
			leastX = newVertex.x
		} else if newVertex.x > mostX {
			mostX = newVertex.x
		}
		if newVertex.y < leastY {
			leastY = newVertex.y
		} else if newVertex.y > mostY {
			mostY = newVertex.y
		}
	}
	fmt.Println(verticies)
	// Build and print the grid
	grid := make([][]Tile, 0)
	for i := leastY; i <= mostY; i++ {
		newRow := make([]Tile, 0)
		for j := leastX; j <= mostX; j++ {
			p := Point{x: j, y: i}
			if PointOnPolygon(p, verticies) {
				newRow = append(newRow, dug)
			} else {
				newRow = append(newRow, undug)
			}
		}
		grid = append(grid, newRow)
	}
	PrintGrid(grid)
	// Flood fill
	output := 0
	filledGrid := FloodFill(grid, Point{x: len(grid[0]) / 2, y: len(grid) / 2})
	fmt.Println()
	PrintGrid(filledGrid)
	for _, line := range filledGrid {
		for _, cell := range line {
			if cell == dug {
				output++
			}
		}
	}
	// Print the result
	fmt.Println(output)
}

func PrintGrid(grid [][]Tile) {
	for _, line := range grid {
		for _, tile := range line {
			fmt.Print(tile)
		}
		fmt.Println()
	}
}

func FloodFill(grid [][]Tile, point Point) [][]Tile {
	if grid[point.y][point.x] == undug {
		grid[point.y][point.x] = dug
		if point.x > 0 {
			FloodFill(grid, Point{x: point.x - 1, y: point.y})
		}
		if point.x < len(grid[0])-1 {
			FloodFill(grid, Point{x: point.x + 1, y: point.y})
		}
		if point.y > 0 {
			FloodFill(grid, Point{x: point.x, y: point.y - 1})
		}
		if point.y < len(grid)-1 {
			FloodFill(grid, Point{x: point.x, y: point.y + 1})
		}
	}
	return grid
}

func PointOnPolygon(point Point, polygon []Point) bool {
	for i := range polygon {
		var a Point
		if i > 0 {
			a = polygon[i-1]
		} else {
			a = polygon[len(polygon)-1]
		}
		b := polygon[i]
		if int(math.Abs(float64(Distance(a, point)+Distance(b, point)-Distance(a, b)))) == 0 {
			return true
		}
	}
	return false
}

func Distance(a Point, b Point) int {
	xDiff := int(math.Abs(float64(b.x - a.x)))
	yDiff := int(math.Abs(float64(b.y - a.y)))
	return xDiff + yDiff
}
