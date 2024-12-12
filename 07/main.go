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

	for _, line := range lines {
		valueToMatch := utils.MustAtoi(strings.Split(line, ": ")[0])
		nums := utils.Map(strings.Split(strings.Split(line, ": ")[1], " "), func(s string, i int) int {
			return utils.MustAtoi(s)
		})

		combinations := allCombinations([]string{"+", "*"}, len(nums)-1)

		for _, combination := range combinations {
			val := nums[0]

			for i := 1; i < len(nums); i++ {
				if combination[i-1] == "+" {
					val += nums[i]
				} else if combination[i-1] == "*" {
					val *= nums[i]
				} else {
					panic("unknown operator")
				}
			}

			if val == valueToMatch {
				result += valueToMatch
				break
			}
		}
	}

	fmt.Println("Part 1 ->", result)
}

func part2(lines []string) {
	var result int

	for _, line := range lines {
		valueToMatch := utils.MustAtoi(strings.Split(line, ": ")[0])
		nums := utils.Map(strings.Split(strings.Split(line, ": ")[1], " "), func(s string, i int) int {
			return utils.MustAtoi(s)
		})

		combinations := allCombinations([]string{"+", "*", "||"}, len(nums)-1)

		for _, combination := range combinations {
			val := nums[0]

			for i := 1; i < len(nums); i++ {
				if combination[i-1] == "+" {
					val += nums[i]
				} else if combination[i-1] == "*" {
					val *= nums[i]
				} else if combination[i-1] == "||" {
					val = utils.MustAtoi(fmt.Sprintf("%d%d", val, nums[i]))
				} else {
					panic("unknown operator")
				}
			}

			if val == valueToMatch {
				result += valueToMatch
				break
			}
		}
	}

	fmt.Println("Part 2 ->", result)
}

func generatePermutations(items []string, ways int, current []string, result *[][]string) {
	// If we have selected 'ways' items, add the current permutation to the result
	if len(current) == ways {
		// Create a copy of the current permutation and add it to the result
		permutation := make([]string, len(current))
		copy(permutation, current)
		*result = append(*result, permutation)
		return
	}

	// Recurse through all items to generate permutations
	for i := 0; i < len(items); i++ {
		// Create a new slice for each recursive call
		newCurrent := make([]string, len(current))
		copy(newCurrent, current)
		// Add the current item to the permutation
		newCurrent = append(newCurrent, items[i])
		// Recurse to fill the next position
		generatePermutations(items, ways, newCurrent, result)
	}
}

// Function to get all combinations of a given size 'ways'
func allCombinations(items []string, ways int) [][]string {
	var result [][]string
	generatePermutations(items, ways, []string{}, &result)
	return result
}
