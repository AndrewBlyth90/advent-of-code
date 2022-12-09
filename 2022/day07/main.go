package main

import (
	"advent-of-code/helpers"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	inputValues := helpers.ReadValuesFromFile("./2022/day07/input.txt")
	partA(inputValues)
}

func partA(inputValues []string){

	path := []string{"/"}
	var finalTotal int64 = 0
	dirCount := 0

	directoryMap := make(map[string]int64)

	for x := 1; x < len(inputValues); x++ {
		if inputValues[x] == "$ cd .." {
			path = path[:len(path)-1]
			continue
		}

		line := strings.Split(inputValues[x], " ")

		if line[0] == "$"{
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

		if line[0] == "dir" {
			dirCount ++
		}
	}

	for _, dir := range directoryMap {
		if dir < 100000 {
			finalTotal += dir
		}
	}

	print("Day 07 Part A: ", finalTotal)
}