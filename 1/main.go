package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("1/input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	left := make([]int, len(lines))
	right := make([]int, len(lines))

	counts_r := make(map[int]int)
	for i, line := range lines {
		splitted := strings.Fields(line)
		if len(splitted) == 0 {
			continue
		}

		if len(splitted) != 2 {
			println(line, len(line), splitted)
			panic("invalid line")
		}
		x, err := strconv.Atoi(splitted[0])
		if err != nil {
			panic(err)
		}
		left[i] = x
		y, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(err)
		}
		right[i] = y

		counts_r[y]++
	}
	slices.Sort(left)
	slices.Sort(right)

	var distance int
	var similarity int
	for i := 0; i < len(left); i++ {
		z := right[i] - left[i]
		if z < 0 {
			z = -z
		}
		distance += z

		similarity += left[i] * counts_r[left[i]]
	}
	println(distance)
	println(similarity)
}
