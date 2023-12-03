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
	gameIdRxp := regexp.MustCompile(`^Game (\d+)$`)
	blueRxp := regexp.MustCompile(`^.*\D(\d+)[\s]blue.*$`)
	redRxp := regexp.MustCompile(`^.*\D(\d+)[\s]red.*$`)
	greenRxp := regexp.MustCompile(`.*\D(\d+)[\s]green.*$`)
	// Get games
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	output := 0
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		headerSplit := strings.Split(line, ":")
		gameId, _ := strconv.Atoi(gameIdRxp.FindStringSubmatch(headerSplit[0])[1])
		reveals := strings.Split(headerSplit[1], ";")
		gameValid := true
		for _, reveal := range reveals {
			// fmt.Println("Reveal ", reveal)
			if !gameValid {
				break
			}
			blueMatch := blueRxp.FindStringSubmatch(reveal)
			// fmt.Println(blueMatch)
			if len(blueMatch) > 1 {
				blue, _ := strconv.Atoi(blueMatch[1])
				// fmt.Println("Blue ", blue)
				if blue > maxBlue {
					// fmt.Println("Excluding game ", gameId, " on reveal ", i, " because blue ", blue)
					gameValid = false
					break
				}
			}
			redMatch := redRxp.FindStringSubmatch(reveal)
			if len(redMatch) > 1 {
				red, _ := strconv.Atoi(redMatch[1])
				// fmt.Println("Red ", red)
				if red > maxRed {
					// fmt.Println("Excluding game ", gameId, " on reveal ", i, " because red ", red)
					gameValid = false
					break
				}
			}
			greenMatch := greenRxp.FindStringSubmatch(reveal)
			if len(greenMatch) > 1 {
				green, _ := strconv.Atoi(greenMatch[1])
				// fmt.Println("Green ", green)
				if green > maxGreen {
					// fmt.Println("Excluding game ", gameId, " on reveal ", i, " because green ", green)
					gameValid = false
					break
				}
			}
		}
		if gameValid {
			fmt.Println("Valid game: ", gameId)
			output += gameId
		}
	}
	// Print the result
	fmt.Println(output)
}
