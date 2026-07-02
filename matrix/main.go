package main

import (
	"fmt"
)

func createMatrix(rows int, columns int) [][]int {
	bigArray := make([][]int, rows)

	for i := 0; i < len(bigArray); i++ {
		subArray := make([]int, columns)
		for j := 0; j < len(subArray); j++ {
			currentElement := i * j
			subArray[j] = currentElement
		}

		bigArray[i] = subArray
	}

	return bigArray
}

func printMatrix(matrix [][] int) {
	for i := 0; i < len(matrix); i++{
		currentRow := matrix[i]
		fmt.Println(currentRow)
	}
}

func main() {
	rows := 10
	columns := 10

	matrix := createMatrix(rows, columns)

	printMatrix(matrix)


	
}
