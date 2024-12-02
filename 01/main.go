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
	var left []int
	var right []int

	for _, line := range lines {
		parts := strings.Fields(line)

		left = append(left, utils.MustAtoi(parts[0]))
		right = append(right, utils.MustAtoi(parts[1]))
	}

	slices.Sort(left)
	slices.Sort(right)

	var result int

	for i, l := range left {
		diff := l - right[i]

		if diff < 0 {
			diff = -diff
		}

		result += diff
	}

	fmt.Println("Part 1 ->", result)
}

func part2(lines []string) {
	var left []int
	right := make(map[int]int)

	for _, line := range lines {
		parts := strings.Fields(line)

		left = append(left, utils.MustAtoi(parts[0]))
		right[utils.MustAtoi(parts[1])]++
	}

	var result int

	for _, v := range left {
		result += v * right[v]
	}

	fmt.Println("Part 2 ->", result)
}
