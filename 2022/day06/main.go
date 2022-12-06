package main

import (
	"advent-of-code/helpers"
)

func unique(strSLice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range strSLice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {

	inputValues := helpers.ReturnStringArrayFromFIle("./2022/day06/input.txt")
	partA(inputValues)
	partB(inputValues)

}

func partA(input []string) {
	for i := 0; i < len(input)-4; i++ {
		uniqueFour := unique(input[i : i+4])
		if len(uniqueFour) == 4 {
			print("Day 06 Part A: ", i+4, "\n")
			return
		}
	}
}

func partB(input []string) {
	for i := 0; i < len(input)-14; i++ {
		uniqueFour := unique(input[i : i+14])
		if len(uniqueFour) == 14 {
			print("Day 06 Part B: ", i+14)
			return
		}
	}
}
