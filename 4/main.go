package main

import (
	"os"
	"slices"
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

	var xmasCounter int
	for lineNum := range grid {
		for colNum := range grid[lineNum] {
			positions := positions(lineNum, colNum)
			for _, position := range positions {
				if checkForXmas(grid, position) {
					xmasCounter++
				}
			}
		}
	}

	println(xmasCounter)
}

func checkForXmas(grid [][]rune, positions [][]int) bool {
	word := []rune("XMAS")
	for i, charPosition := range positions {
		lineNum, colNum := charPosition[0], charPosition[1]
		if lineNum < 0 || colNum < 0 {
			return false
		}
		if lineNum >= len(grid) || colNum >= len(grid[lineNum]) {
			return false
		}

		char := grid[lineNum][colNum]
		if i == 0 && char == word[len(word)-1] {
			// detect reverse word
			slices.Reverse(word)
			continue
		}
		if char != word[i] {
			return false
		}
	}
	return true
}

func positions(lineNum, colNum int) [][][]int {
	return [][][]int{
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
}
