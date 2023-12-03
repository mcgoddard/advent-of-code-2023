package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
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
	lines := strings.Split(content_string, "\n")
	blueRxp := regexp.MustCompile(`^.*\D(\d+)[\s]blue.*$`)
	redRxp := regexp.MustCompile(`^.*\D(\d+)[\s]red.*$`)
	greenRxp := regexp.MustCompile(`.*\D(\d+)[\s]green.*$`)
	// Get games
	output := 0
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		headerSplit := strings.Split(line, ":")
		reveals := strings.Split(headerSplit[1], ";")
		gameRed := 1
		gameGreen := 1
		gameBlue := 1
		for _, reveal := range reveals {
			blueMatch := blueRxp.FindStringSubmatch(reveal)
			if len(blueMatch) > 1 {
				blue, _ := strconv.Atoi(blueMatch[1])
				if blue > gameBlue {
					gameBlue = blue
				}
			}
			redMatch := redRxp.FindStringSubmatch(reveal)
			if len(redMatch) > 1 {
				red, _ := strconv.Atoi(redMatch[1])
				if red > gameRed {
					gameRed = red
				}
			}
			greenMatch := greenRxp.FindStringSubmatch(reveal)
			if len(greenMatch) > 1 {
				green, _ := strconv.Atoi(greenMatch[1])
				if green > gameGreen {
					gameGreen = green
				}
			}
		}
		gamePower := gameBlue * gameGreen * gameRed
		fmt.Println("Game power: ", gamePower)
		output += gamePower
	}
	// Print the result
	fmt.Println(output)
}
