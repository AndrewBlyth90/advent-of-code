package main

import (
	"advent-of-code/helpers"
	"strconv"
	"strings"
)

func buildStacks(input string) []Stack {

	lines := strings.Split(input, "\n")
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	stackCounterArray := strings.Split(lines[0], "")

	var stackArray []Stack

	for _, stackNum := range stackCounterArray {
		if stackNum != " " {
			stackArray = append(stackArray, NewStack())
		}
	}

	for _, stacksLine := range lines[1:] {
		for i, j := 1, 0; i < len(stacksLine); i, j = i+4, j+1 {
			crate := string(stacksLine[i])
			if crate != " " {
				stackArray[j].Push(crate)
			}
		}
	}

	return stackArray

}

func getInputMovements(input string) [][]int {
	inputMovementsLines := strings.Split(input, "\n")
	var movementsArray [][]int

	for _, movementsLine := range inputMovementsLines {

		splitMovements := strings.Split(string(movementsLine), " ")

		if len(splitMovements) != 6 {
			continue
		}

		numberCrates, _ := strconv.Atoi(splitMovements[1])
		initalStack, _ := strconv.Atoi(splitMovements[3])
		finalStack, _ := strconv.Atoi(splitMovements[5])

		movementsArray = append(movementsArray, []int{numberCrates, initalStack - 1, finalStack - 1})
	}

	return movementsArray
}

func main() {

	inputValues := helpers.ReadValuesFromFile("./2022/day05/input.txt")
	stacks := buildStacks(inputValues[0])
	movements := getInputMovements(inputValues[1])
	day01(stacks, movements)

}

func day01(stacks []Stack, movements [][]int) {

	for _, movement := range movements {
		for i := 0; i < movement[0]; i++ {
			if stacks[movement[1]].Empty() {
				break
			}
			crate, _ := stacks[movement[1]].Pop()
			stacks[movement[2]].Push(crate)
		}
	}

	var sb strings.Builder

	for _, stack := range stacks {

		topCrate, _ := stack.Peek()

		sb.WriteString(topCrate.(string))
	}

	print("day05 Part A: ", sb.String())
}
