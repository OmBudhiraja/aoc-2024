package main

import (
	"fmt"
	"slices"

	"github.com/ombudhiraja/aoc-2024/utils"
)

func main() {
	lines := utils.Lines()

	part1(lines[0])
	part2(lines[0])
}

func part1(line string) {

	var result int

	disk := getDisk(line)

	l := 0
	r := len(disk) - 1

	for l < r {
		if disk[l] != "." {
			l++
			continue
		}

		for disk[r] == "." {
			r--
		}

		if l < r {
			disk[l] = disk[r]
			disk[r] = "."

			l++
			r--
		}

	}

	for i, s := range disk {
		if s == "." {
			break
		}

		num := utils.MustAtoi(s)
		result += i * num
	}

	fmt.Println("Part 1 ->", result)
}

func part2(line string) {
	var result int

	disk := getDisk(line)

	for num := len(line) / 2; num >= 0; num-- {
		fileId := fmt.Sprintf("%d", num)
		idStartIdx := slices.Index(disk, fileId)

		if idStartIdx == -1 {
			continue
		}

		fileCount := fileCountWithId(disk, idStartIdx)

		for j := 0; j < idStartIdx; j++ {
			if disk[j] != "." {
				continue
			}

			spaceCount := fileCountWithId(disk, j)

			if spaceCount >= fileCount {
				copy(disk[j:j+fileCount], repeat(fileId, fileCount))
				copy(disk[idStartIdx:idStartIdx+fileCount], repeat(".", fileCount))
				break
			}
		}
	}

	for i, block := range disk {
		if block == "." || block == "_" {
			continue
		}

		result += utils.MustAtoi(block) * i
	}

	fmt.Println("Part 2 ->", result)
}

func getDisk(input string) []string {
	disk := make([]string, 0)

	for i, ch := range input {
		num := utils.MustAtoi(string(ch))

		if i%2 != 0 {
			disk = append(disk, repeat(".", num)...)
		} else {
			id := fmt.Sprintf("%d", i/2)
			disk = append(disk, repeat(id, num)...)
		}

	}

	return disk
}

func fileCountWithId(disk []string, startIndex int) int {
	count := 1

	i := startIndex + 1

	for i < len(disk) && disk[i] == disk[startIndex] {
		count++
		i++
	}

	return count
}

func repeat(val string, count int) []string {
	res := make([]string, count)

	for i := 0; i < count; i++ {
		res[i] = val
	}

	return res
}
