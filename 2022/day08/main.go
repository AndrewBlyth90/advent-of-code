package main

import "advent-of-code/helpers"

type position [2]int

func main() {

	inputValues := helpers.ReadValuesFromFile("./2022/day08/input.txt")
	partA(inputValues)

}


func partA(inputValues []string){
	totalVisible := make(map[position]bool)

	// left to right
	for i, row := range inputValues{
		currentMaxHeight := -1
		for x := 0; x < len(row); x++ {
			if int(row[x]) > currentMaxHeight {
				currentMaxHeight = int(row[x])
				totalVisible[position{x, i}] = true
			}
		}
	}

	// right to left
	for i, row := range inputValues{
		currentMaxHeight := -1
		for x := len(row) - 1; x >= 0; x-- {
			if int(row[x]) > currentMaxHeight {
				currentMaxHeight = int(row[x])
				totalVisible[position{x, i}] = true
			}
		}
	}

	// top to bottom
	for x := 0; x < len(inputValues[0]); x++{
		currentMaxHeight := -1
		for y := 0; y < len(inputValues) - 1; y++{
			if int(inputValues[y][x]) > currentMaxHeight {
				currentMaxHeight = int(inputValues[y][x])
				totalVisible[position{x, y}] = true
			}
		}
	}

	// bottom to top
	for x := len(inputValues[0]) - 1; x > 0; x--{
		currentMaxHeight := -1
		for y := len(inputValues) - 2; y > 0; y--{
			if int(inputValues[y][x]) > currentMaxHeight {
				currentMaxHeight = int(inputValues[y][x])
				totalVisible[position{x, y}] = true
			}
		}
	}

	print("Day 08 Part A: ", len(totalVisible))
}