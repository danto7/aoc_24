package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("2/input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	var safe int
	for _, line := range lines {
		if line == "" {
			continue
		}
		splitted := strings.Fields(line)

		if isSafe(splitted) {
			safe++
		} else {
			newSplitted := make([]string, len(splitted)-1)
			for i := range splitted {
				// safe way to remove element from slice without changing original slice
				copy(newSplitted, splitted[:i])
				copy(newSplitted[i:], splitted[i+1:])

				if isSafe(newSplitted) {
					safe++
					break
				}
			}
		}
	}
	println(safe)
}

func isSafe(splitted []string) bool {
	var increase bool
	var last_num int
	for i := 0; i < len(splitted); i++ {
		n, err := strconv.Atoi(splitted[i])
		if err != nil {
			panic(err)
		}

		if i == 0 {
			// skip first number
			last_num = n
			continue
		} else if i == 1 {
			// set icrease or decrease flag
			if last_num > n {
				increase = false
			} else {
				increase = true
			}
		}

		if increase {
			if n <= last_num {
				return false
			}
		} else {
			if n >= last_num {
				return false
			}
		}

		distance := abs(n - last_num)
		if distance < 1 || distance > 3 {
			return false
		}
		last_num = n
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
