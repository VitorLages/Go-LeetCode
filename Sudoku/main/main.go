package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sudoku, err := os.ReadFile("sudoku.txt")

	if err != nil {
		fmt.Println("Error Reading File: ", err)
	}

	sudokuMatrix := formatTile(string(sudoku))
	resolve(sudokuMatrix, 0)
}

func resolve(sudoku [][]string, pivot int) {
	var line []string
	var column []string
	var square []string

	if pivot == 9 {
		return
	}

	for i := range sudoku {
		for j := range sudoku {

			if i == pivot {
				if sudoku[i][j] != "0" {
					line = append(line, sudoku[i][j])
				}
			}

			if (i != pivot && j == pivot) || (i == pivot && j == pivot) {
				if sudoku[i][j] != "0" {
					column = append(column, sudoku[i][j])
				}
			}

			if i == 0 && j == 0 {
				for k := range 3 {
					for h := range 3 {
						if sudoku[k][h] != "0" {
							square = append(square, sudoku[k][h])
						}
					}
				}
			}
		}
	}
	fmt.Println("line: ", line)
	fmt.Println("column: ", column)
	fmt.Println("square: ", square)
	resolve(sudoku, pivot+1)
}

func formatTile(sudoku string) [][]string {
	sudokuArray := []rune(sudoku)
	var sudokuString string
	offset := 2
	skip := 35

	for i := range 9 {
		slice := sudokuArray[offset+(i*skip) : offset+skip+(i*skip)]
		temp := strings.Replace(string(slice), `"`, "", -1)
		temp = strings.Replace(temp, `.`, "0", -1)
		temp = strings.Replace(temp, `,`, "", -1)
		sudokuString += temp
		offset += 3
	}

	sudokuNewArray := []rune(sudokuString)
	count := 0

	sudokuMatrix := make([][]string, 9)
	for i := range sudokuMatrix {
		sudokuMatrix[i] = make([]string, 9)
	}

	for i := range 9 {
		for k := range 9 {
			sudokuMatrix[i][k] = string(sudokuNewArray[count])
			count++
		}
	}
	return sudokuMatrix
}
