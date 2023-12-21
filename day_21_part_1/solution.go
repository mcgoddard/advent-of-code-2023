package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Tile int

const (
	start  Tile = iota
	garden Tile = iota
	rock   Tile = iota
)

var tileMap = map[rune]Tile{
	'S': start,
	'.': garden,
	'#': rock,
}

type Point struct {
	x int
	y int
}

type Node struct {
	point Point
	steps int
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
	grid := make([][]Tile, len(lines))
	var startPoint Node
	for i, line := range lines {
		row := make([]Tile, len([]rune(line)))
		for j, char := range line {
			row[j] = tileMap[char]
			if row[j] == start {
				startPoint.point = Point{x: j, y: i}
			}
		}
		grid[i] = row
	}
	// Build graph
	queue := []Node{startPoint}
	visited := make(map[Point]int)
	neighbours := []Point{
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: 0, y: 1},
	}
	for len(queue) > 0 {
		next := queue[0]
		if next.steps > 6 {
			break
		}
		queue = queue[1:]
		for _, neighbour := range neighbours {
			neighbourPoint := Point{x: next.point.x + neighbour.x, y: next.point.y + neighbour.y}
			if _, exists := visited[neighbourPoint]; neighbourPoint.x >= 0 && neighbourPoint.x < len(grid[0]) &&
				neighbourPoint.y >= 0 && neighbourPoint.y < len(grid) &&
				grid[neighbourPoint.y][neighbourPoint.x] == garden && !exists {
				queue = append(queue, Node{point: neighbourPoint, steps: next.steps + 1})
				visited[neighbourPoint] = next.steps + 1
			}
		}
	}
	// Count visited points with steps % 2 == 0
	output := 1
	for _, s := range visited {
		if s%2 == 0 {
			output++
		}
	}
	// Print the result
	fmt.Println(output)
}
