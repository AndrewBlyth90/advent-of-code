package main

import "advent-of-code/helpers"

func main() {
	inputValues := helpers.ReadValuesFromFile("./2022/day02/input.txt")
	partA(inputValues)
	partB(inputValues)
}

func partA(inputValues []string) {
	scoreMap := map[string]int{"A X": 4, "B X": 1, "C X": 7, "A Y": 8, "B Y": 5, "C Y": 2, "A Z": 3, "B Z": 9, "C Z": 6}
	total := 0
	for _, item := range inputValues {
		total = total + scoreMap[item]
	}
	println("day02 part A: ", total)
}

func partB(inputValues []string) {
	scoreMap := map[string]int{"A X": 3, "B X": 1, "C X": 2, "A Y": 4, "B Y": 5, "C Y": 6, "A Z": 8, "B Z": 9, "C Z": 7}
	total := 0
	for _, item := range inputValues {
		total = total + scoreMap[item]
	}
	println("day02 part B: ", total)
}
