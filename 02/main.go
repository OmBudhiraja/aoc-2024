package main

import (
	"fmt"
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

	for _, report := range lines {
		levels := utils.Map(strings.Fields(report), func(val string, index int) int {
			return utils.MustAtoi(val)
		})

		if isSafe(levels) {
			result++
		}
	}

	fmt.Println("Part 1 ->", result)
}

func part2(lines []string) {
	var result int

	for _, report := range lines {
		levels := utils.Map(strings.Fields(report), func(val string, index int) int {
			return utils.MustAtoi(val)
		})

		if isSafe(levels) {
			result++
			continue
		}

		for i := 0; i < len(levels); i++ {
			clone := make([]int, len(levels))
			copy(clone, levels)

			if isSafe(append(clone[:i], clone[i+1:]...)) {
				result++
				break
			}

		}
	}

	fmt.Println("Part 2 ->", result)
}

func isSafe(levels []int) bool {

	isIncreasing := levels[1] > levels[0]

	for i := 1; i < len(levels); i++ {
		delta := levels[i] - levels[i-1]

		if isIncreasing {
			if delta <= 0 || delta > 3 {
				return false
			}
		} else {
			if delta >= 0 || delta < -3 {
				return false
			}
		}
	}

	return true
}
