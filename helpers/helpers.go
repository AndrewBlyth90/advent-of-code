package helpers

import (
	"fmt"
	"os"
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
	result := 0
	for _, number := range numbers {
		result *= number
	}
	return result
}
