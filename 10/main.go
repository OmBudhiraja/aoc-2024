package main

import (
	"fmt"
	"strings"

	"github.com/ombudhiraja/aoc-2024/utils"
)

type Point struct {
	x int
	y int
}

var directions [][]int = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func main() {
	lines := utils.Lines()
	grid := utils.Map(lines, func(line string, i int) []int {
		return utils.Map(strings.Split(line, ""), func(c string, i int) int {
			return utils.MustAtoi(c)
		})
	})

	part1(grid)
	part2(grid)
}

func part1(grid [][]int) {
	trailheadPositions := []Point{}
	endPositions := []Point{}

	var result int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				trailheadPositions = append(trailheadPositions, Point{
					x: j,
					y: i,
				})
			}

			if grid[i][j] == 9 {
				endPositions = append(endPositions, Point{
					x: j,
					y: i,
				})
			}
		}
	}

	for _, trailhead := range trailheadPositions {
		for _, endPos := range endPositions {
			if distinctPaths(grid, trailhead, endPos, -1) > 0 {
				result++
			}
		}
	}

	fmt.Println("Part 1 ->", result)
}

func part2(grid [][]int) {
	trailheadPositions := []Point{}
	endPositions := []Point{}

	var result int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				trailheadPositions = append(trailheadPositions, Point{
					x: j,
					y: i,
				})
			}

			if grid[i][j] == 9 {
				endPositions = append(endPositions, Point{
					x: j,
					y: i,
				})
			}
		}
	}

	for _, trailhead := range trailheadPositions {
		for _, endPos := range endPositions {
			result += distinctPaths(grid, trailhead, endPos, -1)
		}
	}

	fmt.Println("Part 2 ->", result)
}

func distinctPaths(grid [][]int, cur, end Point, lastNum int) int {

	if cur.y < 0 || cur.y >= len(grid) || cur.x < 0 || cur.x >= len(grid[cur.y]) {
		return 0
	}

	if grid[cur.y][cur.x] != lastNum+1 {
		return 0
	}

	if cur.y == end.y && cur.x == end.x {
		return 1
	}

	paths := 0

	for _, dir := range directions {
		paths += distinctPaths(grid, Point{
			x: cur.x + dir[0],
			y: cur.y + dir[1],
		}, end, lastNum+1)
	}

	return paths
}
