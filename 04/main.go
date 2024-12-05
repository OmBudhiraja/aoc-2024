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
	grid := utils.Map(lines, func(line string, index int) []string {
		return strings.Split(line, "")
	})

	grid = padGrid(grid, 3)

	var result int

	pointsForAllDir := func(i, j int) []string {
		return []string{
			grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i][j+3],       // forward
			grid[i][j] + grid[i][j-1] + grid[i][j-2] + grid[i][j-3],       // backward
			grid[i][j] + grid[i+1][j] + grid[i+2][j] + grid[i+3][j],       // downward
			grid[i][j] + grid[i-1][j] + grid[i-2][j] + grid[i-3][j],       // upward
			grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2] + grid[i+3][j+3], // downward right diagonal
			grid[i][j] + grid[i+1][j-1] + grid[i+2][j-2] + grid[i+3][j-3], // downward left diagonal
			grid[i][j] + grid[i-1][j-1] + grid[i-2][j-2] + grid[i-3][j-3], // upward right diagonal
			grid[i][j] + grid[i-1][j+1] + grid[i-2][j+2] + grid[i-3][j+3], // upward left diagonal
		}
	}

	for i := 3; i < len(grid)-3; i++ {
		for j := 3; j < len(grid[i])-3; j++ {
			if grid[i][j] != "X" {
				continue
			}

			matches := len(utils.Filter(pointsForAllDir(i, j), func(word string, index int) bool {
				return word == "XMAS"
			}))

			result += matches
		}
	}

	fmt.Println("Part 1 ->", result)
}

func part2(lines []string) {
	grid := utils.Map(lines, func(line string, index int) []string {
		return strings.Split(line, "")
	})

	grid = padGrid(grid, 2)

	var result int

	for i := 2; i < len(grid)-2; i++ {
		for j := 2; j < len(grid[i])-2; j++ {
			if grid[i][j] != "M" && grid[i][j] != "S" {
				continue
			}

			firstW := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
			secondW := grid[i+2][j] + grid[i+1][j+1] + grid[i][j+2]

			if (firstW == "MAS" || firstW == "SAM") && (secondW == "MAS" || secondW == "SAM") {
				result++
			}
		}
	}

	fmt.Println("Part 2 ->", result)
}

func padGrid(grid [][]string, paddingSize int) [][]string {
	rows := len(grid)
	cols := len(grid[0])

	paddedGrid := make([][]string, rows+2*paddingSize)
	for i := range paddedGrid {
		paddedGrid[i] = make([]string, cols+2*paddingSize)
	}

	// Fill the padded grid with '.' by default
	for i := range paddedGrid {
		for j := range paddedGrid[i] {
			paddedGrid[i][j] = "."
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			paddedGrid[i+paddingSize][j+paddingSize] = grid[i][j]
		}
	}

	return paddedGrid
}
