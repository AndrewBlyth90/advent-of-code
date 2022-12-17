package main

import (
	"advent-of-code/helpers"
	"golang.org/x/exp/maps"
	"strconv"
	"strings"
)

func main(){
	inputValues := helpers.ReadValuesFromFileDoubleSpace("./2022/day11/input.txt")
	monkeyMap := createMonkeyMap(inputValues)
	monkeyMapTotalItems := make(map[int]int)


	for x := 1; x <= 20; x++ {
		for i, instruction := range inputValues{
			newWorryMultiplactor, divisableBy, ifTrue, ifFalse, operation := getValues(instruction)

			monkeyMapTotalItems[i] = monkeyMapTotalItems[i] + len(monkeyMap[i])

			for _, item := range monkeyMap[i]{
				var multiply int

				if newWorryMultiplactor == "old" {
					multiply = item
				} else {
					multiply, _ = strconv.Atoi(newWorryMultiplactor)
				}

				newWorryLevel := getWorryLevel(item, multiply, operation)

				if newWorryLevel % divisableBy == 0 {
					monkeyMap[ifTrue] = append(monkeyMap[ifTrue], newWorryLevel)
				}
				if newWorryLevel % divisableBy > 0 {
					monkeyMap[ifFalse] = append(monkeyMap[ifFalse], newWorryLevel)
				}
			}
			monkeyMap[i] = []int{}
		}
	}
	monkeyBusiness := getMonkeyBusiness(monkeyMapTotalItems)
	print("Day11 Part A: ", monkeyBusiness)
}

func getMonkeyBusiness(items map[int]int) int{
	monkeyBusinessSlice := maps.Values(items)
	monkeyBusinessSlice = helpers.SortSliceDecending(monkeyBusinessSlice)
	return monkeyBusinessSlice[0] * monkeyBusinessSlice[1]
}


func getWorryLevel (item int, multiply int, operation string) int {

	if operation == "*"{
		return (item * multiply) / 3
	}

	return (item + multiply) / 3
}

func getValues (instruction string) (string, int, int, int, string){
	splitValues := strings.Split(instruction, "\n")
	newWorryMultiplactor := strings.Split(splitValues[2], " ")
	divisableBy := strings.Split(splitValues[3], "by ")
	ifTrue := strings.Split(splitValues[4], "monkey ")
	ifFalse := strings.Split(splitValues[5], "monkey ")

	divisibleByInt, _ := strconv.Atoi(divisableBy[1])
	ifTrueInt, _ := strconv.Atoi(ifTrue[1])
	ifFalseInt, _ := strconv.Atoi(ifFalse[1])

	return newWorryMultiplactor[7], divisibleByInt, ifTrueInt, ifFalseInt, newWorryMultiplactor[6]
}

func createMonkeyMap (inputValues []string) map[int][]int {
	monkeyMap := make(map[int][]int)

	for i, instruction := range inputValues{
		splitValues := strings.Split(instruction, "\n")
		items := strings.Split(splitValues[1], ": ")
		individualItems := strings.Split(items[1], ", ")
		itemsAsInts := []int{}
		for _, individualItem := range individualItems {
			int, _ := strconv.Atoi(individualItem)
			itemsAsInts = append(itemsAsInts, int)
		}
		monkeyMap[i] = itemsAsInts
	}
	return monkeyMap
}
