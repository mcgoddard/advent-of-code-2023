package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type TileType int

const (
	northToSouth TileType = iota
	eastToWest   TileType = iota
	northToEast  TileType = iota
	northToWest  TileType = iota
	southToWest  TileType = iota
	southToEast  TileType = iota
	ground       TileType = iota
	start        TileType = iota
	unset        TileType = iota
)

var tileMap = map[string]TileType{
	"|": northToSouth,
	"-": eastToWest,
	"L": northToEast,
	"J": northToWest,
	"7": southToWest,
	"F": southToEast,
	".": ground,
	"S": start,
}

type Point struct {
	x int
	y int
}

var connections = map[TileType][]Point{
	northToSouth: {
		{x: 0, y: -1},
		{x: 0, y: 1},
	},
	eastToWest: {
		{x: -1, y: 0},
		{x: 1, y: 0},
	},
	northToEast: {
		{x: 0, y: -1},
		{x: 1, y: 0},
	},
	northToWest: {
		{x: 0, y: -1},
		{x: -1, y: 0},
	},
	southToWest: {
		{x: 0, y: 1},
		{x: -1, y: 0},
	},
	southToEast: {
		{x: 0, y: 1},
		{x: 1, y: 0},
	},
	ground: {},
	start: {
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: 0, y: 1},
	},
}

type void struct{}

var member void

func main() {
	exampleFlagPtr := flag.Bool("example", false, "Use the example file if set")
	example2FlagPtr := flag.Bool("example2", false, "Use the second example file if set")
	example3FlagPtr := flag.Bool("example3", false, "Use the third example file if set")
	flag.Parse()
	// Read input file
	inputFile := "input.txt"
	if *exampleFlagPtr {
		inputFile = "example.txt"
	} else if *example2FlagPtr {
		inputFile = "example2.txt"
	} else if *example3FlagPtr {
		inputFile = "example3.txt"
	}
	content, _ := os.ReadFile(inputFile)
	content_string := string(content)
	// Split on newline
	lines := strings.Split(content_string, "\n")
	grid := make([][]TileType, 0)
	// Read map
	startPosition := Point{x: 0, y: 0}
	for i, line := range lines {
		cells := make([]TileType, 0)
		for j, cell := range line {
			cellType := tileMap[string(cell)]
			if cellType == start {
				startPosition = Point{x: j, y: i}
			}
			cells = append(cells, cellType)
		}
		grid = append(grid, cells)
	}
	// Search from start
	visitedPoints := make(map[Point]void)
	route := []Point{startPosition}
	for len(route) < 2 || !reflect.DeepEqual(route[len(route)-1], startPosition) {
		currentPosition := route[len(route)-1]
		currentTile := grid[currentPosition.y][currentPosition.x]
		currentConnections := connections[currentTile]
		for _, connection := range currentConnections {
			// Get the next position if this connection were to be followed
			nextPosition := Point{
				x: currentPosition.x + connection.x,
				y: currentPosition.y + connection.y,
			}
			// If it's out of bounds then skip
			if nextPosition.x < 0 || nextPosition.x >= len(grid[0]) || nextPosition.y < 0 || nextPosition.y >= len(grid) {
				continue
			}
			// If it would backtrack then skip
			if len(route) > 1 {
				lastPosition := route[len(route)-2]
				if lastPosition.x == nextPosition.x && lastPosition.y == nextPosition.y {
					continue
				}
			}
			// Otherwise check it's not a dead end
			nextTile := grid[nextPosition.y][nextPosition.x]
			nextConnections := connections[nextTile]
			validConnection := false
			for _, nextConnection := range nextConnections {
				if nextConnection.x*-1 == connection.x && nextConnection.y*-1 == connection.y {
					validConnection = true
					break
				}
			}
			// Add to the route and iterate
			if validConnection {
				route = append(route, nextPosition)
				visitedPoints[nextPosition] = member
				break
			}
		}
	}
	// Work out what S is and replace with it's actual character
	firstPosition := route[1]
	lastPosition := route[len(route)-2]
	for tile, connections := range connections {
		if tile == start {
			continue
		}
		firstDiff := Point{firstPosition.x - startPosition.x, firstPosition.y - startPosition.y}
		secondDiff := Point{lastPosition.x - startPosition.x, lastPosition.y - startPosition.y}
		match := true
		for _, connection := range connections {
			if !reflect.DeepEqual(firstDiff, connection) && !reflect.DeepEqual(secondDiff, connection) {
				match = false
			}
		}
		if match {
			grid[startPosition.y][startPosition.x] = tile
			break
		}
	}
	// Walk the grid counting inside vs outside
	output := 0
	for i := range grid {
		out := true
		for j := range grid[0] {
			currentPoint := Point{x: j, y: i}
			if _, exists := visitedPoints[currentPoint]; exists {
				if grid[i][j] == northToSouth || grid[i][j] == northToWest || grid[i][j] == northToEast {
					out = !out
				}
			} else if !out {
				output++
			}
		}
	}
	// Print the result
	fmt.Println(output)
}
