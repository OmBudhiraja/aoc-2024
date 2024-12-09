package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ombudhiraja/aoc-2024/utils"
)

func main() {
	lines := utils.Lines()

	part1(lines)
	part2(lines)
}

type Guard struct {
	x   int
	y   int
	dir [2]int
}

type Point struct {
	x int
	y int
}

func part1(lines []string) {
	grid := utils.Map(lines, func(line string, index int) []string {
		return strings.Split(line, "")
	})

	visitedCoords := map[Point]bool{}

	guard := getGuard(grid)
	visitedCoords[Point{x: guard.x, y: guard.y}] = true

	for guard.x < len(grid[0])-1 && guard.x > 0 && guard.y < len(grid)-1 && guard.y > 0 {
		guard.move(grid)
		visitedCoords[Point{x: guard.x, y: guard.y}] = true
	}

	fmt.Println("Part 1->", len(visitedCoords))
}

func part2(lines []string) {
	grid := utils.Map(lines, func(line string, index int) []string {
		return strings.Split(line, "")
	})

	var result int

	guard := getGuard(grid)

	startPoint := Point{x: guard.x, y: guard.y}
	startDir := guard.dir

	path := map[Point]bool{}

	for guard.x < len(grid[0])-1 && guard.x > 0 && guard.y < len(grid)-1 && guard.y > 0 {
		guard.move(grid)
		path[Point{x: guard.x, y: guard.y}] = true
	}

	for point := range path {
		if point.x == startPoint.x && point.y == startPoint.y {
			continue
		}

		if grid[point.y][point.x] != "." {
			continue
		}

		grid[point.y][point.x] = "#"

		if isLoopCreated(grid, Guard{
			x:   startPoint.x,
			y:   startPoint.y,
			dir: startDir,
		}) {
			result++
		}

		grid[point.y][point.x] = "."
	}

	fmt.Println("Part 2->", result)

}

func getGuard(grid [][]string) *Guard {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			cur := grid[i][j]
			if cur == "^" || cur == ">" || cur == "v" || cur == "<" {

				dir := [2]int{}

				if cur == "^" {
					dir[0] = 0
					dir[1] = -1
				}

				if cur == ">" {
					dir[0] = 1
					dir[1] = 0
				}

				if cur == "<" {
					dir[0] = -1
					dir[1] = 0
				}

				if cur == "v" {
					dir[0] = 0
					dir[1] = 1
				}

				return &Guard{
					x:   j,
					y:   i,
					dir: dir,
				}
			}
		}
	}

	panic("no guard found")
}

func inGridBounds(grid [][]string, x, y int) bool {
	return x < len(grid[0]) && x >= 0 && y < len(grid) && y >= 0
}

func (g *Guard) move(grid [][]string) {
	nextX := g.x + g.dir[0]
	nextY := g.y + g.dir[1]

	if !inGridBounds(grid, nextX, nextY) {
		g.x = nextX
		g.y = nextY
		return
	}

	// if obstacle move 90 degrees right
	if grid[nextY][nextX] == "#" {
		// g.dir = g.getNextDir()

		tmp := g.dir[0]
		if g.dir[0] == 0 {
			g.dir[0] = -1 * g.dir[1]
		} else {
			g.dir[0] = 0
		}

		if g.dir[1] == 0 {
			g.dir[1] = tmp
		} else {
			g.dir[1] = 0
		}

		return
	}

	g.x = nextX
	g.y = nextY
}

func isLoopCreated(grid [][]string, guard Guard) bool {

	visitedStates := map[string]bool{}

	for inGridBounds(grid, guard.x, guard.y) {
		stateKey := generateKey(guard.y, guard.x, guard.dir[0], guard.dir[1])

		if _, exists := visitedStates[stateKey]; exists {
			return true
		}

		visitedStates[stateKey] = true

		guard.move(grid)
	}

	return false
}

func generateKey(nums ...int) string {
	s := utils.Map(nums, func(n int, i int) string {
		return strconv.Itoa(n)
	})

	return strings.Join(s, ".")
}
