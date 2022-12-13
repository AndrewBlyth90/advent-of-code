package main

import (
	"advent-of-code/helpers"
	"strconv"
	"strings"
)

func main() {

	inputValues := helpers.ReadValuesFromFile("./2022/day09/input.txt")

	var head, tail [2]int

	directions := map[string][2]int{
		"U": {1, 0},
		"D": {-1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}

	visted := map[[2]int]bool{
		{0, 0}: true,
	}

	for _, instruction := range inputValues {
		splitInstructions := strings.Split(instruction, " ")

		for x, _ := strconv.Atoi(splitInstructions[1]); x > 0; x-- {
			directionToMove := directions[splitInstructions[0]]

			head[0] += directionToMove[0]
			head[1] += directionToMove[1]

			rowDifference := head[0] - tail[0]
			colDifference := head[1] - tail[1]

			if helpers.AbsInt(rowDifference) > 1 {
				tail[0] += rowDifference / 2

				if colDifference != 0 {
					tail[1] += colDifference
				}
			} else if helpers.AbsInt(colDifference) > 1 {
				tail[1] += colDifference / 2
				if rowDifference != 0 {
					tail[0] += rowDifference
				}
			}

			visted[tail] = true
		}

	}
	print(len(visted))
}
