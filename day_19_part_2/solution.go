package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	x Range
	m Range
	a Range
	s Range
}

type Range struct {
	start int
	end   int
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
	// Run parts through workflows
	part := Part{
		x: Range{start: 1, end: 4000},
		m: Range{start: 1, end: 4000},
		a: Range{start: 1, end: 4000},
		s: Range{start: 1, end: 4000},
	}
	output := AcceptPart(part, workflows, "in")
	// Print output
	fmt.Println(output)
}

func AcceptPart(part Part, workflows map[string]Workflow, currentWorkflow string) int {
	// Check if we've reached an accept or reject
	if currentWorkflow == "A" {
		return (part.x.end - part.x.start + 1) * (part.m.end - part.m.start + 1) * (part.a.end - part.a.start + 1) * (part.s.end - part.s.start + 1)
	} else if currentWorkflow == "R" {
		return 0
	}
	// Otherwise apply rules
	result := 0
	for _, rule := range workflows[currentWorkflow].rules {
		// Have we reached a default apply?
		if rule.category == "" {
			result += AcceptPart(part, workflows, rule.outcome)
		}
		// Test rule
		newPart := part
		greaterThan := rule.operator == ">"
		if greaterThan {
			if rule.category == "x" {
				newPart.x.start = rule.operand + 1
				part.x.end = rule.operand
				result += AcceptPart(newPart, workflows, rule.outcome)
			} else if rule.category == "m" {
				newPart.m.start = rule.operand + 1
				part.m.end = rule.operand
				result += AcceptPart(newPart, workflows, rule.outcome)
			} else if rule.category == "a" {
				newPart.a.start = rule.operand + 1
				part.a.end = rule.operand
				result += AcceptPart(newPart, workflows, rule.outcome)
			} else if rule.category == "s" {
				newPart.s.start = rule.operand + 1
				part.s.end = rule.operand
				result += AcceptPart(newPart, workflows, rule.outcome)
			}
		} else {
			if rule.category == "x" {
				newPart.x.end = rule.operand - 1
				part.x.start = rule.operand
				result += AcceptPart(newPart, workflows, rule.outcome)
			} else if rule.category == "m" {
				newPart.m.end = rule.operand - 1
				part.m.start = rule.operand
				result += AcceptPart(newPart, workflows, rule.outcome)
			} else if rule.category == "a" {
				newPart.a.end = rule.operand - 1
				part.a.start = rule.operand
				result += AcceptPart(newPart, workflows, rule.outcome)
			} else if rule.category == "s" {
				newPart.s.end = rule.operand - 1
				part.s.start = rule.operand
				result += AcceptPart(newPart, workflows, rule.outcome)
			}
		}
	}
	return result
}
