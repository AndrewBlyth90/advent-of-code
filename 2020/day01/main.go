package main

import (
	"advent-of-code/helpers"
)

func main() {

	inputValues := helpers.ReadValuesFromFile("./2020/day01/input.txt")
	inputInts := helpers.ConvertToInts(inputValues)

	for _, number := range inputInts {
		remaining := 2020 - number
		if helpers.Contains(inputInts, remaining) {
			println(number * remaining)
		}
	}

}
