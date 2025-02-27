package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var output string
	fmt.Print("Enter Numbers For First Array: ")
	fmt.Scan(&input)
	l1 := input

	if len(input) != 0 {
		l1 = invert(l1)
	}

	fmt.Print("Enter Numbers For First Array: ")
	fmt.Scan(&input)
	l2 := input

	if len(input) != 0 {
		l2 = invert(l2)
	}

	output = sumArrays(l1, l2)
	output = invert(output)
	fmt.Println("Your sum is: ", output)
}

func invert(arr string) string {
	newArray := []rune(arr)
	invertArray := make([]rune, len(newArray))

	for i, pos := range newArray {
		invertArray[len(newArray)-1-i] = pos
	}

	return string(invertArray)
}

func sumArrays(l1 string, l2 string) string {
	newl1, _ := strconv.ParseInt(l1, 10, 0)
	newl2, _ := strconv.ParseInt(l2, 10, 0)
	output := newl1 + newl2
	return strconv.Itoa(int(output))
}
