package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type ParserMode int

const (
	RuleParsing ParserMode = iota
	PageOrderParsing
)

type Rule struct {
	Left  int
	Right int
}
type Rules []Rule

type PageOrder []int

func main() {
	data, err := os.ReadFile("5/input")
	if err != nil {
		panic(err)
	}
	var rules Rules
	var pageOrders []PageOrder

	parsingMode := RuleParsing

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			parsingMode = PageOrderParsing
			continue
		}
		switch parsingMode {
		case RuleParsing:
			parts := strings.Split(line, "|")
			left, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			rules = append(rules, Rule{Left: left, Right: right})
		case PageOrderParsing:
			parts := strings.Split(line, ",")
			var pageOrder PageOrder
			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err != nil {
					panic(err)
				}
				pageOrder = append(pageOrder, num)
			}
			pageOrders = append(pageOrders, pageOrder)
		}
	}

	var wrongPageOrders []PageOrder
	var accCorrectPageOrder int
	for _, pageOrder := range pageOrders {
		if checkRules(rules, pageOrder) {
			if len(pageOrder)%2 != 1 {
				panic("page order length must be odd")
			}
			accCorrectPageOrder += pageOrder[len(pageOrder)/2]
		} else {
			wrongPageOrders = append(wrongPageOrders, pageOrder)
		}
	}

	fmt.Println("part one:", accCorrectPageOrder)

	/*
		resultChannel := make(chan int, 10)
		var wg sync.WaitGroup
		wg.Add(len(wrongPageOrders))
		for _, pageOrder := range wrongPageOrders {
			pageOrder := pageOrder // create a new local variable
			go func() {
				defer wg.Done()
				resultChannel <- findCorrectPageOrder(pageOrder, rules)
			}()
		}
		go func() {
			wg.Wait()
			close(resultChannel)
		}()

		var reorderedPageOrdersAcc int
		for num := range resultChannel {
			reorderedPageOrdersAcc += num
		}
	*/

	fmt.Println("part two:", reorderedPageOrdersAcc)
}

func findCorrectPageOrder(pageOrder PageOrder, rules Rules) int {
	var applicableRules Rules
	for _, rule := range rules {
		if slices.Contains(pageOrder, rule.Left) && slices.Contains(pageOrder, rule.Right) {
			applicableRules = append(applicableRules, rule)
		}
	}

	permutations := permutate(pageOrder)
	for _, permutation := range permutations {
		if checkRules(applicableRules, permutation) {
			fmt.Println("old page order", pageOrder)
			fmt.Println("new page order", permutation)
			return permutation[len(permutation)/2]
		}
	}
	panic("no correct page order found")
}
func permutate(pageOrder PageOrder) []PageOrder {
	var pageOrders []PageOrder
	if len(pageOrder) == 1 {
		return []PageOrder{pageOrder}
	}
	if len(pageOrder) == 2 {
		return []PageOrder{pageOrder, {pageOrder[1], pageOrder[0]}}
	}
	for i, num := range pageOrder {
		rest := make(PageOrder, len(pageOrder)-1)
		copy(rest, pageOrder[:i])
		copy(rest[i:], pageOrder[i+1:])
		restPermutations := permutate(rest)
		for _, restPermutation := range restPermutations {
			pageOrders = append(pageOrders, append([]int{num}, restPermutation...))
		}
	}

	return pageOrders
}

func checkRules(rules Rules, pageOrder PageOrder) bool {
	for _, rule := range rules {
		if !checkRule(rule, pageOrder) {
			return false
		}
	}
	return true
}

func checkRule(rule Rule, pageOrder PageOrder) bool {
	leftSideFound := false
	rightSideFound := false
	for _, page := range pageOrder {
		switch page {
		case rule.Left:
			leftSideFound = true

			// right side was found before left side
			if rightSideFound {
				return false
			}
		case rule.Right:
			rightSideFound = true

			// left side was already found
			if leftSideFound {
				return true
			}
		}
	}
	// one of the side or both were not found
	// rule is satisfied
	return true
}
