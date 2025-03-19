package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error Reading File: ", err)
	}

	numbers, target := format(string(file))
	tempTarget := []rune(target)
	intTarget, _ := strconv.Atoi(string(tempTarget[0]))

	intNumbers := stringToIntArray(numbers)

	position := twoSums(intNumbers, intTarget)
	fmt.Println(position)
}

func twoSums(numbers []int, target int) []int {
	prev := make(map[int]int)
	var position []int

	for i := range numbers {
		diff := target - numbers[i]

		if _, ok := prev[diff]; ok {
			position = append(position, prev[diff], i)
		}

		prev[numbers[i]] = i
	}

	return position
}

func stringToIntArray(numbers []string) []int {
	var intArray []int
	for i := range numbers {
		aux, _ := strconv.Atoi(numbers[i])
		intArray = append(intArray, aux)
	}
	return intArray
}

func format(file string) ([]string, string) {
	getTarget := strings.LastIndex(string(file), "=")

	if getTarget == -1 {
		fmt.Println("Target not found")
	}

	target := file[getTarget+2:] // +2 because "= " the equal symbol and the space right after

	input := strings.FieldsFunc(string(file), func(r rune) bool {
		if r == '[' || r == ']' {
			return true
		}
		return false
	})

	aux := input[1]
	inputArray := []rune(aux)

	numbersArray := strings.FieldsFunc(string(inputArray), func(r rune) bool {
		if r == ',' {
			return true
		}
		return false
	})

	return numbersArray, target
}
