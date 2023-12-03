package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
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
	content_string := string(content)
	// Split on newline
	lines := strings.Split(content_string, "\n")
	// First and last number in each line
	firstNumberRxp := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	lastNumberRxp := regexp.MustCompile(`(\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)
	numberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	numbers := []int{}
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		firstNumber := firstNumberRxp.FindStringSubmatch(line)[1]
		if _, err := strconv.Atoi(firstNumber); err != nil {
			firstNumber = numberMap[firstNumber]
		}
		reverseLine := reverse(line)
		lastNumber := lastNumberRxp.FindStringSubmatch(reverseLine)[1]
		if _, err := strconv.Atoi(lastNumber); err != nil {
			lastNumber = reverse(lastNumber)
			lastNumber = numberMap[lastNumber]
		}
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
