package main

import (
	"advent-of-code/helpers"
	"strconv"
	"strings"
)

func main() {
	inputValues := helpers.ReadValuesFromFile("./2022/day04/input.txt")
	partA(inputValues)
	partB(inputValues)

}

func partA(inputValues []string) {
	total := 0
	for _, item := range inputValues {
		numbers := splitNumbers(item)
		if numbers[2] >= numbers[0] && numbers[3] <= numbers[1] || numbers[2] <= numbers[0] && numbers[3] >= numbers[1] {
			total = total + 1
		}
	}
	print("day04 part A: ", total, "\n")
}

func partB(inputValues []string) {
	total := 0
	for _, item := range inputValues {
		numbers := splitNumbers(item)

		if numbers[0] >= numbers[2] &&
			numbers[0] <= numbers[3] ||
			numbers[2] >= numbers[0] &&
				numbers[2] <= numbers[1] {
			total = total + 1
		}
	}
	print("day04 part B: ", total)
}

func splitNumbers(values string) []int {
	split := strings.Split(values, ",")
	var numbers []int
	for _, pair := range split {
		individualNumber := strings.Split(pair, "-")
		for _, number := range individualNumber {
			intNumber, _ := strconv.Atoi(number)
			numbers = append(numbers, intNumber)
		}
	}
	return numbers
}
