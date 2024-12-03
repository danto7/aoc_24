package main

import (
	"os"
	"regexp"
	"strconv"
)

var partOne = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func main() {
	text, err := os.ReadFile("3/input")
	if err != nil {
		panic(err)
	}
	doPartOne(text)
}

func doPartOne(text []byte) {
	matches := partOne.FindAllStringSubmatch(string(text), -1)
	var acc int64
	for _, match := range matches {
		a, err := strconv.ParseInt(match[1], 10, 64)
		if err != nil {
			panic(err)
		}
		b, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil {
			panic(err)
		}
		acc += a * b
	}
	println(acc)
}
