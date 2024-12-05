package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ombudhiraja/aoc-2024/utils"
)

func main() {
	lines := utils.Lines()

	part1(lines)
	part2(lines)
}

func part1(lines []string) {

	var result int

	rules := make(map[int][]int)
	var rulesEndIdx int

	for i, line := range lines {
		if line == "" {
			rulesEndIdx = i + 1
			break
		}

		nums := strings.Split(line, "|")
		x := utils.MustAtoi(nums[0])
		y := utils.MustAtoi(nums[1])

		rules[x] = append(rules[x], y)
	}

	for i := rulesEndIdx; i < len(lines); i++ {
		updates := utils.Map(strings.Split(lines[i], ","), func(s string, index int) int {
			return utils.MustAtoi(s)
		})

		valid := isCorrectOrder(rules, updates)

		if valid {
			middle := updates[len(updates)/2]
			result += middle
		}
	}

	fmt.Println("Part 1 ->", result)
}

func part2(lines []string) {
	var result int

	rules := make(map[int][]int)
	var rulesEndIdx int

	for i, line := range lines {
		if line == "" {
			rulesEndIdx = i + 1
			break
		}

		nums := strings.Split(line, "|")
		x := utils.MustAtoi(nums[0])
		y := utils.MustAtoi(nums[1])

		rules[x] = append(rules[x], y)
	}

	for i := rulesEndIdx; i < len(lines); i++ {
		updates := utils.Map(strings.Split(lines[i], ","), func(s string, index int) int {
			return utils.MustAtoi(s)
		})

		valid := isCorrectOrder(rules, updates)

		if valid {
			continue
		}

		for i := len(updates) - 1; i >= 0; i-- {
			for j, before := range updates[0:i] {
				if slices.Contains(rules[updates[i]], before) {
					// swap
					tmp := updates[j]
					updates[j] = updates[i]
					updates[i] = tmp
				}
			}
		}

		result += updates[len(updates)/2]

	}

	fmt.Println("Part 2 ->", result)
}

func isCorrectOrder(rules map[int][]int, updates []int) bool {
	for i, num := range updates {
		for _, after := range updates[i+1:] {
			rule := rules[after]

			if slices.Contains(rule, num) {
				return false
			}
		}

	}

	return true
}
