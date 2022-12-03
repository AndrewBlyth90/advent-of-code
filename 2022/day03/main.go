package main

import (
	"advent-of-code/helpers"
	"strings"
)

// take the input
// split it in half
// find out which item appears in both halves
// find the value of that item

var itemMap = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52}

func main() {
	inputValues := helpers.ReadValuesFromFile("./2022/day03/input.txt")
	partA(inputValues)
	partB(inputValues)
}

func partA(inputValues []string) {
	total := 0

	for _, item := range inputValues {
		var commonLetter string
		halfwayPoint := len(item) / 2
		firstHalf := item[:halfwayPoint]
		secondHalf := item[halfwayPoint:]
		for _, letter := range firstHalf {
			if strings.Contains(secondHalf, string(letter)) {
				commonLetter = string(letter)
			}
		}
		total = total + itemMap[commonLetter]
	}
	println("Day 03 part A: ", total)
}

func partB(inputValues []string) {

	var result [][]string
	total := 0

	for i := 0; i < len(inputValues)/3; i++ {
		min := i * 3
		max := (i + 1) * 3

		result = append(result, inputValues[min:max])
	}

	for _, trilogy := range result {
		for _, line := range trilogy {
			for _, letter := range line {
				if strings.Contains(trilogy[1], string(letter)) {
					if strings.Contains(trilogy[2], string(letter)) {
						total = total + itemMap[string(letter)]
						break
					}
				}
			}
			break
		}
	}
	println("Day 03 part B: ", total)
}
