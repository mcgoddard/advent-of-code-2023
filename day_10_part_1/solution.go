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

func main() {
	exampleFlagPtr := flag.Bool("example", false, "Use the example file if set")
	example2FlagPtr := flag.Bool("example2", false, "Use the second example file if set")
	flag.Parse()
	// Read input file
	inputFile := "input.txt"
	if *exampleFlagPtr {
		inputFile = "example.txt"
	} else if *example2FlagPtr {
		inputFile = "example2.txt"
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
				break
			}
		}
	}
	// Print the result
	fmt.Println(len(route) / 2)
}
