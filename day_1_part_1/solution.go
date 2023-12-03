package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read input file
	content, _ := os.ReadFile("input.txt")
	content_string := string(content)
	// Split on newline
	lines := strings.Split(content_string, "\n")
	// First and last number in each line
	firstNumberRxp := regexp.MustCompile(`(\d)`)
	lastNumberRxp := regexp.MustCompile(`^.*(\d)([a-zA-Z])*$`)
	numbers := []int{}
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		firstNumber := firstNumberRxp.FindStringSubmatch(line)[1]
		lastNumber := lastNumberRxp.FindStringSubmatch(line)[1]
		i, _ := strconv.Atoi(firstNumber + lastNumber)
		numbers = append(numbers, i)
	}
	// Add numbers
	output := 0
	for _, number := range numbers {
		output += number
	}
	// Print the result
	fmt.Println(output)
}
