package main

import (
	"advent-of-code/helpers"
	"strconv"
	"strings"
)

func main() {

	inputValues := helpers.ReadValuesFromFile("./2022/day10/input.txt")

	cycles := []int{20, 60, 100, 140, 180, 220}
	total := make([]int, 0)
	numberOfCycles := 0
	registerTotal := 1


	for _, instruction := range inputValues {

		splitInstructions := strings.Split(instruction, " ")
		if len(splitInstructions) == 1 {
			numberOfCycles ++
			if helpers.Contains(cycles, numberOfCycles){
				total = append(total, registerTotal * numberOfCycles)
			}
			continue
		}



		for x := 0; x < 3; x++ {
			numberOfCycles++
			if helpers.Contains(cycles, numberOfCycles){
				total = append(total, registerTotal * numberOfCycles)
			}
		}

		numberOfCycles++
		if helpers.Contains(cycles, numberOfCycles){
			total = append(total, registerTotal * numberOfCycles)
		}

		registerIncrease, _ := strconv.Atoi(splitInstructions[1])
		registerTotal = registerIncrease + registerTotal

	}
	print("Day10 part A: ", helpers.SumArray(total...))
}


