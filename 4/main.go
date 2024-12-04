package main

import (
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("4/input")
	if err != nil {
		panic(err)
	}
	var grid [][]rune
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		grid = append(grid, []rune(line))
	}

	var partOneCounter int
	var partTwoCounter int
	for lineNum := range grid {
		for colNum := range grid[lineNum] {
			// PART TWO
			wordPositions := [][][]int{
				{
					// diagonal left down
					{lineNum, colNum},
					{lineNum + 1, colNum - 1},
					{lineNum + 2, colNum - 2},
					{lineNum + 3, colNum - 3},
				}, {
					// diagonal right down
					{lineNum, colNum},
					{lineNum + 1, colNum + 1},
					{lineNum + 2, colNum + 2},
					{lineNum + 3, colNum + 3},
				}, {
					// horizontal right
					{lineNum, colNum},
					{lineNum, colNum + 1},
					{lineNum, colNum + 2},
					{lineNum, colNum + 3},
				}, {
					// vertical down
					{lineNum, colNum},
					{lineNum + 1, colNum},
					{lineNum + 2, colNum},
					{lineNum + 3, colNum},
				},
			}
			for _, position := range wordPositions {
				if checkForWordInPositions(grid, "XMAS", position) || checkForWordInPositions(grid, "SAMX", position) {
					partOneCounter++
				}
			}

			// PART ONE

			// 1 . 2
			// . 3 .
			// 4 . 5
			wordPosition := [][]int{
				{lineNum, colNum},
				{lineNum, colNum + 2},
				{lineNum + 1, colNum + 1},
				{lineNum + 2, colNum},
				{lineNum + 2, colNum + 2},
			}

			words := []string{
				// M . S
				// . A .
				// M . S
				"MSAMS",

				// S . M
				// . A .
				// S . M
				"SMASM",

				// S . S
				// . A .
				// M . M
				"SSAMM",

				// M . M
				// . A .
				// S . S
				"MMASS",
			}

			for _, word := range words {
				if checkForWordInPositions(grid, word, wordPosition) {
					partTwoCounter++
				}
			}
		}
	}

	println("partone: ", partOneCounter)
	println("parttwo: ", partTwoCounter)
}

func isWithinGridBounds(grid [][]rune, lineNum, colNum int) bool {
	return lineNum >= 0 && lineNum < len(grid) && colNum >= 0 && colNum < len(grid[lineNum])
}

func checkForWordInPositions(grid [][]rune, word string, positions [][]int) bool {
	chars := []rune(word)
	for i, charPosition := range positions {
		lineNum, colNum := charPosition[0], charPosition[1]
		if !isWithinGridBounds(grid, lineNum, colNum) {
			return false
		}
		char := grid[lineNum][colNum]

		if char != chars[i] {
			return false
		}
	}
	return true
}
