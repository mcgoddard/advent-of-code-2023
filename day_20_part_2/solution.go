package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type ModuleType int

const (
	flipFlop    ModuleType = iota
	conjunction ModuleType = iota
	broadcaster ModuleType = iota
)

type Module struct {
	moduleType       ModuleType
	targets          []string
	name             string
	flipFlopState    bool
	conjunctionState map[string]bool
}

type Pulse struct {
	source      string
	destination string
	high        bool
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
	modules := make(map[string]Module)
	for _, line := range lines {
		newModule := Module{}
		firstChar := []rune(line)[0]
		parts := strings.Split(string([]rune(line)[1:]), " -> ")
		newModule.name = parts[0]
		newModule.targets = strings.Split(parts[1], ", ")
		if firstChar == '%' {
			newModule.moduleType = flipFlop
		} else if firstChar == '&' {
			newModule.moduleType = conjunction
			newModule.conjunctionState = make(map[string]bool)
		} else {
			newModule.moduleType = broadcaster
			newModule.name = "broadcaster"
		}
		modules[newModule.name] = newModule
	}
	// Set inputs for cons
	for key, module := range modules {
		if module.moduleType == conjunction {
			for _, sourceModule := range modules {
				for _, target := range sourceModule.targets {
					if target == key {
						module.conjunctionState[sourceModule.name] = false
						modules[key] = module
						break
					}
				}
			}
		}
	}
	// Identify the feed modules for conjunction module that pulses to rx (this only works because one conjunction feeds rx)
	feedsRx := ""
	for _, module := range modules {
		for _, target := range module.targets {
			if target == "rx" {
				feedsRx = module.name
				break
			}
		}
		if feedsRx != "" {
			break
		}
	}
	lookingFor := make(map[string]int)
	for _, module := range modules {
		for _, target := range module.targets {
			if target == feedsRx {
				lookingFor[module.name] = -1
				break
			}
		}
	}
	// Run pulses
	buttonPresses := 0
	pulseDelivered := false
	for !pulseDelivered {
		buttonPresses++
		pulses := []Pulse{{
			source:      "button",
			destination: "broadcaster",
			high:        false,
		}}
		for len(pulses) > 0 {
			// Pull the first pulse off
			currentPulse := pulses[0]
			// Check if we've found a module we're looking for
			for lookingForTarget, value := range lookingFor {
				if value < 0 && currentPulse.high && currentPulse.source == lookingForTarget {
					lookingFor[lookingForTarget] = buttonPresses
					allFound := true
					for _, value := range lookingFor {
						if value < 0 {
							allFound = false
							break
						}
					}
					if allFound {
						pulseDelivered = true
						break
					}
				}
			}
			pulses = pulses[1:]
			// Process and retreive further pulses
			newPulses := processPulse(modules, currentPulse)
			// Append new pulses to the end
			pulses = append(pulses, newPulses...)
		}
	}
	// Lowest common multiple of each feed modules cycle
	output := 1
	for _, value := range lookingFor {
		output = (output * value) / greatestCommonDivisor(output, value)
	}
	// Print the result
	fmt.Println(output)
}

func greatestCommonDivisor(a, b int) int {
	for b > 0 {
		c := b
		b = a % b
		a = c
	}
	return a
}

func processPulse(modules map[string]Module, pulse Pulse) []Pulse {
	newPulses := make([]Pulse, 0)
	targetModule := modules[pulse.destination]
	if targetModule.moduleType == broadcaster {
		for _, target := range targetModule.targets {
			newPulses = append(newPulses, Pulse{
				source:      targetModule.name,
				destination: target,
				high:        pulse.high,
			})
		}
	} else if targetModule.moduleType == flipFlop {
		if !pulse.high {
			targetModule.flipFlopState = !targetModule.flipFlopState
			modules[pulse.destination] = targetModule
			for _, target := range targetModule.targets {
				newPulses = append(newPulses, Pulse{
					source:      targetModule.name,
					destination: target,
					high:        targetModule.flipFlopState,
				})
			}
		}
	} else if targetModule.moduleType == conjunction {
		targetModule.conjunctionState[pulse.source] = pulse.high
		modules[pulse.destination] = targetModule
		allHigh := true
		for _, value := range targetModule.conjunctionState {
			if !value {
				allHigh = false
			}
		}
		for _, target := range targetModule.targets {
			newPulses = append(newPulses, Pulse{
				source:      targetModule.name,
				destination: target,
				high:        !allHigh,
			})
		}
	}
	return newPulses
}
