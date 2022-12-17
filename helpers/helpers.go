package helpers

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadValuesFromFile(filepath string) []string {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	return strings.Split(string(contents), "\n")
}

func ReadValuesFromFileDoubleSpace(filepath string) []string {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	return strings.Split(string(contents), "\n\n")
}

func AbsInt(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func ReturnStringArrayFromFIle(filepath string) []string {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	return strings.Split(string(contents), "")
}

func ConvertToInts(strArr []string) (values []int) {
	for _, val := range strArr {
		val, _ := strconv.Atoi(val)
		values = append(values, val)
	}
	return values
}

func Contains(intArray []int, singleInt int) bool {

	for _, v := range intArray {
		if v == singleInt {
			return true
		}
	}
	return false
}

func MultiplyArray(numbers ...int) int {
	result := 1
	for _, number := range numbers {
		result *= number
	}
	return result
}

func SumArray(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SortSliceDecending(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return slice
}

func RemoveFromSlice(s []int, i int) []int {
	if len(s) == 1{
		return []int{}
	}
	s[i] = s[0]
	return s[1:]
}


