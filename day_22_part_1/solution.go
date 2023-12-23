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
	z int
}

type Block struct {
	start Point
	end   Point
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
	blocks := make(map[int]Block)
	maxX := 0
	maxY := 0
	maxZ := 0
	for i, line := range lines {
		points := strings.Split(line, "~")
		newBlock := Block{}
		pointParts := strings.Split(points[0], ",")
		newBlock.start.x, _ = strconv.Atoi(pointParts[0])
		newBlock.start.y, _ = strconv.Atoi(pointParts[1])
		newBlock.start.z, _ = strconv.Atoi(pointParts[2])
		pointParts = strings.Split(points[1], ",")
		newBlock.end.x, _ = strconv.Atoi(pointParts[0])
		newBlock.end.y, _ = strconv.Atoi(pointParts[1])
		newBlock.end.z, _ = strconv.Atoi(pointParts[2])
		blocks[i+1] = newBlock
		if newBlock.start.x > maxX {
			maxX = newBlock.start.x
		}
		if newBlock.end.x > maxX {
			maxX = newBlock.end.x
		}
		if newBlock.start.y > maxY {
			maxY = newBlock.start.y
		}
		if newBlock.end.y > maxY {
			maxY = newBlock.end.y
		}
		if newBlock.start.z > maxZ {
			maxZ = newBlock.start.z
		}
		if newBlock.end.z > maxZ {
			maxZ = newBlock.end.z
		}
	}
	// Build grid
	grid := make([][][]int, maxX+1)
	for i := 0; i <= maxX; i++ {
		grid[i] = make([][]int, maxY+1)
		for j := 0; j <= maxY; j++ {
			grid[i][j] = make([]int, maxZ+1)
		}
	}
	for id, block := range blocks {
		for i := block.start.x; i <= block.end.x; i++ {
			for j := block.start.y; j <= block.end.y; j++ {
				for k := block.start.z; k <= block.end.z; k++ {
					grid[i][j][k] = id
				}
			}
		}
	}
	// Drop blocks
	blockMoved := true
	for blockMoved {
		blockMoved = false
		for id, block := range blocks {
			newZ := 1
			for i := block.start.x; i <= block.end.x; i++ {
				for j := block.start.y; j <= block.end.y; j++ {
					minZ := block.start.z
					if block.end.z < block.start.z {
						newZ = block.end.z
					}
					for k := minZ - 1; k >= 1; k-- {
						if grid[i][j][k] == 0 {
							minZ = k
						} else {
							break
						}
					}
					if minZ > newZ {
						newZ = minZ
					}
				}
			}
			if newZ < block.start.z && newZ < block.end.z {
				blockMoved = true
				blocks[id] = Block{
					start: Point{x: block.start.x, y: block.start.y, z: newZ},
					end:   Point{x: block.end.x, y: block.end.y, z: newZ + block.end.z - block.start.z},
				}
				for i := block.start.x; i <= block.end.x; i++ {
					for j := block.start.y; j <= block.end.y; j++ {
						for k := block.start.z; k <= block.end.z; k++ {
							grid[i][j][k] = 0
						}
					}
				}
				for i := blocks[id].start.x; i <= blocks[id].end.x; i++ {
					for j := blocks[id].start.y; j <= blocks[id].end.y; j++ {
						for k := blocks[id].start.z; k <= blocks[id].end.z; k++ {
							grid[i][j][k] = id
						}
					}
				}
			}
		}
	}
	// Find supported and supports for each block
	supports := make(map[int][]int)
	supported := make(map[int][]int)
	for id, block := range blocks {
		supports[id] = make([]int, 0)
		supported[id] = make([]int, 0)
		for i := block.start.x; i <= block.end.x; i++ {
			for j := block.start.y; j <= block.end.y; j++ {
				maxZ := int(math.Max(float64(block.start.z), float64(block.end.z)))
				if grid[i][j][maxZ+1] != 0 && !NumberInSlice(supports[id], grid[i][j][maxZ+1]) {
					supports[id] = append(supports[id], grid[i][j][maxZ+1])
				}
				minZ := int(math.Max(float64(block.start.z), float64(block.end.z)))
				if grid[i][j][minZ-1] != 0 && !NumberInSlice(supported[id], grid[i][j][minZ-1]) {
					supported[id] = append(supported[id], grid[i][j][minZ-1])
				}
			}
		}
	}
	// Find blocks that can safely be destroyed
	output := 0
	for id := range blocks {
		destroyable := true
		for _, supportingId := range supports[id] {
			if len(supported[supportingId]) < 2 {
				destroyable = false
				break
			}
		}
		if destroyable {
			output++
		}
	}
	// Print the result
	fmt.Println(output)
}

func NumberInSlice(slice []int, number int) bool {
	for _, item := range slice {
		if item == number {
			return true
		}
	}
	return false
}
