package main

import (
	"advent-of-code/helpers"
	"sort"
)

func main() {

	inputValues := helpers.ReadValuesFromFile("./2022/day01/input.txt")
	inputInts := helpers.ConvertToInts(inputValues)

	partA(inputInts)
	partB(inputInts)

}

func partA(inputInts []int) {

	highest := 0
	runningTotal := 0

	for _, number := range inputInts {
		if number == 0 {
			if runningTotal > highest {
				highest = runningTotal
			}
			runningTotal = 0
		}
		runningTotal = runningTotal + number
	}

	println("day01 part A: ", highest)
}

func partB(inputInts []int) {
	runningTotal := 0
	var amounts []int

	for _, number := range inputInts {
		if number == 0 {
			amounts = append(amounts, runningTotal)
			runningTotal = 0
		}
		runningTotal = runningTotal + number
	}

	sort.Slice(amounts, func(i, j int) bool {
		return amounts[i] > amounts[j]
	})

	println("day01 part A: ", amounts[0]+amounts[1]+amounts[2])

}
