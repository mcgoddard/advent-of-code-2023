package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Tile int

const (
	empty             Tile = iota
	leftRightMirror   Tile = iota
	rightLeftMirror   Tile = iota
	horizonalSplitter Tile = iota
	verticalSplitter  Tile = iota
)

var tileMap = map[rune]Tile{
	'.':  empty,
	'\\': leftRightMirror,
	'/':  rightLeftMirror,
	'-':  horizonalSplitter,
	'|':  verticalSplitter,
}

type Point struct {
	x int
	y int
}

type Vector struct {
	x int
	y int
}

type Ray struct {
	point  Point
	vector Vector
}

type void struct{}

var member void

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
	for i, line := range lines {
		grid[i] = make([]Tile, len([]rune(line)))
		for j, char := range line {
			grid[i][j] = tileMap[char]
		}
	}
	// Follow rays
	startPoint := Point{x: 0, y: 0}
	energised := map[Point]int{
		startPoint: 1,
	}
	seen := make(map[Ray]void)
	rays := []Ray{{
		point: startPoint,
		vector: Vector{
			x: 1,
			y: 0,
		},
	}}
	for len(rays) > 0 {
		ray := rays[0]
		newRays := moveRay(grid, ray, energised, seen)
		rays = append(rays[1:], newRays...)
	}
	// Print the result
	fmt.Println(len(energised))
}

func moveRay(grid [][]Tile, ray Ray, energised map[Point]int, seen map[Ray]void) []Ray {
	// Bump energised count
	energised[ray.point] += 1
	// Record ray as seen
	newRays := make([]Ray, 0)
	if _, exists := seen[ray]; exists {
		return newRays
	} else {
		seen[ray] = member
	}
	// Calculate new ray(s)
	currentTile := grid[ray.point.y][ray.point.x]
	if currentTile == empty || (currentTile == horizonalSplitter && ray.vector.y == 0) || (currentTile == verticalSplitter && ray.vector.x == 0) {
		newPoint := Point{
			x: ray.point.x + ray.vector.x,
			y: ray.point.y + ray.vector.y,
		}
		if insideGrid(grid, newPoint) {
			newRays = append(newRays, Ray{
				point:  newPoint,
				vector: ray.vector,
			})
		}
	} else if currentTile == leftRightMirror {
		newVector := Vector{
			x: ray.vector.y,
			y: ray.vector.x,
		}
		newPoint := Point{
			x: ray.point.x + newVector.x,
			y: ray.point.y + newVector.y,
		}
		if insideGrid(grid, newPoint) {
			newRays = append(newRays, Ray{
				point:  newPoint,
				vector: newVector,
			})
		}
	} else if currentTile == rightLeftMirror {
		newVector := Vector{
			x: ray.vector.y * -1,
			y: ray.vector.x * -1,
		}
		newPoint := Point{
			x: ray.point.x + newVector.x,
			y: ray.point.y + newVector.y,
		}
		if insideGrid(grid, newPoint) {
			newRays = append(newRays, Ray{
				point:  newPoint,
				vector: newVector,
			})
		}
	} else if currentTile == horizonalSplitter {
		newRay := Ray{
			vector: Vector{
				x: -1,
				y: 0,
			},
			point: Point{
				x: ray.point.x - 1,
				y: ray.point.y,
			},
		}
		if insideGrid(grid, newRay.point) {
			newRays = append(newRays, newRay)
		}
		newRay = Ray{
			vector: Vector{
				x: 1,
				y: 0,
			},
			point: Point{
				x: ray.point.x + 1,
				y: ray.point.y,
			},
		}
		if insideGrid(grid, newRay.point) {
			newRays = append(newRays, newRay)
		}
	} else if currentTile == verticalSplitter {
		newRay := Ray{
			vector: Vector{
				x: 0,
				y: -1,
			},
			point: Point{
				x: ray.point.x,
				y: ray.point.y - 1,
			},
		}
		if insideGrid(grid, newRay.point) {
			newRays = append(newRays, newRay)
		}
		newRay = Ray{
			vector: Vector{
				x: 0,
				y: 1,
			},
			point: Point{
				x: ray.point.x,
				y: ray.point.y + 1,
			},
		}
		if insideGrid(grid, newRay.point) {
			newRays = append(newRays, newRay)
		}
	}
	return newRays
}

func insideGrid(grid [][]Tile, point Point) bool {
	return point.x >= 0 && point.x < len(grid[0]) && point.y >= 0 && point.y < len(grid)
}
