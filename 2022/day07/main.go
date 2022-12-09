package main

import (
	"advent-of-code/helpers"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputValues := helpers.ReadValuesFromFile("./2022/day07/input.txt")
	directoryMap := createDirectoryMap(inputValues)

	partA(directoryMap)
	partB(directoryMap)
}

func createDirectoryMap(inputValues []string) map[string]int64 {
	path := []string{"/"}
	directoryMap := make(map[string]int64)

	for x := 1; x < len(inputValues); x++ {
		if inputValues[x] == "$ cd .." {
			path = path[:len(path)-1]
			continue
		}

		line := strings.Split(inputValues[x], " ")

		if line[0] == "$" {
			if line[1] == "cd" {
				rename := line[2]
				if _, ok := directoryMap[rename]; ok {
					rename = line[2] + strconv.Itoa(rand.Int())
				}
				path = append(path, rename)
			}
			continue
		}

		if line[0] != "dir" {
			size, _ := strconv.ParseInt(line[0], 10, 64)
			for _, p := range path {
				directoryMap[p] += size
			}
		}
	}

	return directoryMap

}

func partA(directoryMap map[string]int64) {
	var finalTotal int64 = 0

	for _, dir := range directoryMap {
		if dir < 100000 {
			finalTotal += dir
		}
	}

	print("Day 07 Part A: ", finalTotal, "\n")
}

func partB(directoryMap map[string]int64) {
	remainingFileSize := 70000000 - directoryMap["/"]
	sizeNeeded := 30000000 - remainingFileSize

	var validDirectory = make([]int64, 0)

	for _, value := range directoryMap {
		if value >= sizeNeeded {
			validDirectory = append(validDirectory, value)
		}
	}

	sort.Slice(validDirectory, func(i, j int) bool {
		return validDirectory[i] < validDirectory[j]
	})

	print("day 07 part B: ", validDirectory[0])
}
