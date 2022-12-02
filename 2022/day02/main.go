package main

import "advent-of-code/helpers"

// A = Rock - 1pt
// B = Paper - 2pt
// C = Scissors - 3pt

// X = Rock - 1pt
// Y = Paper - 2pt
// Z = Scissors - 3pt

// A X = 0pt + 3pt = 3pt
// B X = 0pt + 1pt = 1pt
// C X = 0pt + 2pt = 2pt

// A Y = 3pt + 1pt = 4pt
// B Y = 3pt + 2pt = 5pt
// C Y = 3pt + 3pt = 6pt

// A Z = 6pt + 2pt = 8pt
// B Z = 6pt + 3pt = 9pt
// C Z = 6pt + 1pt = 7pt

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
