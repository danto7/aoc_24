package main

import (
	"os"
	"regexp"
	"strconv"
)

var partOne = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var partTwo = regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)

func main() {
	text, err := os.ReadFile("3/input")
	if err != nil {
		panic(err)
	}
	println("Part one:")
	doPartOne(text)
	println("Part two:")
	doPartTwo(text)
}

func doPartTwo(text []byte) {
	matches := partTwo.FindAllStringSubmatch(string(text), -1)
	var acc int64
	skip := false
	for _, match := range matches {
		if match[0] == "do()" {
			skip = false
			continue
		} else if match[0] == "don't()" {
			skip = true
			continue
		}
		if skip {
			continue
		}

		a, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil {
			panic(err)
		}
		b, err := strconv.ParseInt(match[3], 10, 64)
		if err != nil {
			panic(err)
		}
		acc += a * b
	}
	println(acc)
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
