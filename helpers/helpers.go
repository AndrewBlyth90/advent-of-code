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
