package main

import (
	"fmt"
	"strings"

	"github.com/ombudhiraja/aoc-2024/utils"
)

func main() {
	lines := utils.Lines()

	grid := utils.Map(lines, func(line string, i int) []string {
		return strings.Split(line, "")
	})

	part1(grid)
	part2(grid)

}

type Point struct {
	y int
	x int
}

func part1(grid [][]string) {
	frequencyPoints := map[string][]Point{}

	uniqueAnitinodes := map[Point]bool{}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "." {
				continue

			}

			frequencyPoints[grid[i][j]] = append(frequencyPoints[grid[i][j]], Point{
				y: i,
				x: j,
			})
		}
	}

	for _, freq := range frequencyPoints {
		for i := 0; i < len(freq); i++ {
			for j := i + 1; j < len(freq); j++ {
				xDiff := freq[i].x - freq[j].x
				yDiff := freq[i].y - freq[j].y

				if xDiff < 0 {
					xDiff = -xDiff
				}

				if yDiff < 0 {
					yDiff = -yDiff
				}

				var antinodeA Point
				var antinodeB Point

				antinodeA.y = freq[i].y - yDiff
				antinodeB.y = freq[j].y + yDiff

				if freq[i].x > freq[j].x {
					antinodeA.x = freq[i].x + xDiff
					antinodeB.x = freq[j].x - xDiff
				} else {
					antinodeA.x = freq[i].x - xDiff
					antinodeB.x = freq[j].x + xDiff
				}

				if inGridBounds(grid, antinodeA) {
					uniqueAnitinodes[antinodeA] = true
				}

				if inGridBounds(grid, antinodeB) {
					uniqueAnitinodes[antinodeB] = true
				}
			}
		}
	}

	fmt.Println(len(uniqueAnitinodes))
}

func part2(grid [][]string) {
	frequencyPoints := map[string][]Point{}

	uniqueAnitinodes := map[Point]bool{}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "." {
				continue

			}

			frequencyPoints[grid[i][j]] = append(frequencyPoints[grid[i][j]], Point{
				y: i,
				x: j,
			})
		}
	}

	for _, freq := range frequencyPoints {
		for i := 0; i < len(freq); i++ {
			for j := i + 1; j < len(freq); j++ {
				xDiff := freq[i].x - freq[j].x
				yDiff := freq[i].y - freq[j].y

				if xDiff < 0 {
					xDiff = -xDiff
				}

				if yDiff < 0 {
					yDiff = -yDiff
				}

				uniqueAnitinodes[freq[i]] = true
				uniqueAnitinodes[freq[j]] = true

				n := 1

				// all up antinodes
				for {
					antinode := Point{}
					antinode.y = freq[i].y - (n * yDiff)

					if freq[i].x > freq[j].x {
						antinode.x = freq[i].x + (n * xDiff)
					} else {
						antinode.x = freq[i].x - (n * xDiff)
					}

					if inGridBounds(grid, antinode) {
						uniqueAnitinodes[antinode] = true
					} else {
						break
					}

					n++
				}

				n = 1
				// all down antinodes
				for {
					antinode := Point{}
					antinode.y = freq[j].y + (n * yDiff)

					if freq[i].x > freq[j].x {
						antinode.x = freq[j].x - (n * xDiff)
					} else {
						antinode.x = freq[j].x + (n * xDiff)
					}

					if inGridBounds(grid, antinode) {
						uniqueAnitinodes[antinode] = true
					} else {
						break
					}

					n++
				}
			}
		}
	}

	fmt.Println(len(uniqueAnitinodes))
}

func inGridBounds(grid [][]string, p Point) bool {
	return p.x < len(grid[0]) && p.x >= 0 && p.y < len(grid) && p.y >= 0
}
