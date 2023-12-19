package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Workflow struct {
	name  string
	rules []Rule
}

type Rule struct {
	category string
	operator string
	operand  int
	outcome  string
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
	initialSplit := strings.Split(contentString, "\n\n")
	// Extract workflows
	workflowStrings := strings.Split(initialSplit[0], "\n")
	workflows := make(map[string]Workflow)
	for _, workflowString := range workflowStrings {
		workFlowStringInitialSplit := strings.Split(workflowString, "{")
		name := workFlowStringInitialSplit[0]
		workFlowStringInitialSplit[1] = strings.Replace(workFlowStringInitialSplit[1], "}", "", -1)
		steps := strings.Split(workFlowStringInitialSplit[1], ",")
		rules := make([]Rule, 0)
		for _, r := range steps[:len(steps)-1] {
			category := string([]rune(r)[0])
			operator := string([]rune(r)[1])
			operandOutcomeSplit := strings.Split(string([]rune(r)[2:]), ":")
			operand, _ := strconv.Atoi(operandOutcomeSplit[0])
			outcome := operandOutcomeSplit[1]
			rules = append(rules, Rule{
				category: category,
				operator: operator,
				operand:  operand,
				outcome:  outcome,
			})
		}
		rules = append(rules, Rule{
			category: "",
			operator: "",
			operand:  0,
			outcome:  steps[len(steps)-1],
		})
		workflows[name] = Workflow{
			name:  name,
			rules: rules,
		}
	}
	// Extract parts
	partStrings := strings.Split(initialSplit[1], "\n")
	parts := make([]Part, 0)
	for _, partString := range partStrings {
		partString := strings.Replace(partString, "{", "", -1)
		partString = strings.Replace(partString, "}", "", -1)
		partStringSections := strings.Split(partString, ",")
		partValuesMap := make(map[string]int)
		for _, section := range partStringSections {
			split := strings.Split(section, "=")
			value, _ := strconv.Atoi(split[1])
			partValuesMap[split[0]] = value
		}
		part := Part{
			x: partValuesMap["x"],
			m: partValuesMap["m"],
			a: partValuesMap["a"],
			s: partValuesMap["s"],
		}
		parts = append(parts, part)
	}
	// Run parts through workflows
	accepted := make([]Part, 0)
	for _, part := range parts {
		if AcceptPart(part, workflows, "in") {
			accepted = append(accepted, part)
		}
	}
	// Print the result
	output := 0
	for _, part := range accepted {
		output += part.x + part.m + part.a + part.s
	}
	fmt.Println(output)
}

func AcceptPart(part Part, workflows map[string]Workflow, currentWorkflow string) bool {
	// Check if we've reached an accept or reject
	if currentWorkflow == "A" {
		return true
	} else if currentWorkflow == "R" {
		return false
	}
	// Otherwise apply rules
	for _, rule := range workflows[currentWorkflow].rules {
		// Have we reached a default apply?
		if rule.category == "" {
			return AcceptPart(part, workflows, rule.outcome)
		}
		// Test rule
		greaterThan := rule.operator == ">"
		partValue := 0
		if rule.category == "x" {
			partValue = part.x
		} else if rule.category == "m" {
			partValue = part.m
		} else if rule.category == "a" {
			partValue = part.a
		} else if rule.category == "s" {
			partValue = part.s
		}
		if (greaterThan && partValue > rule.operand) || (!greaterThan && partValue < rule.operand) {
			return AcceptPart(part, workflows, rule.outcome)
		}
	}
	return false
}
