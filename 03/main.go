package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"

	"github.com/ombudhiraja/aoc-2024/utils"
)

func main() {
	lines := utils.Lines()

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	reader := bufio.NewReader(strings.NewReader(strings.Join(lines, "")))

	var result int

	for {
		startBytes, err := reader.ReadBytes(byte('('))

		if errors.Is(err, io.EOF) {
			break
		} else {
			utils.CheckError(err)
		}

		if !strings.HasSuffix(string(startBytes), "mul(") {
			continue
		}

		endBytes, err := reader.ReadBytes(byte(')'))

		if errors.Is(err, io.EOF) {
			break
		} else {
			utils.CheckError(err)
		}

		endBytes = endBytes[:len(endBytes)-1]

		var data string
		var garbage string

		for _, b := range endBytes {
			if unicode.IsDigit(rune(b)) || b == ',' {
				data += string(b)
			} else {
				data = ""
				garbage += string(b)
			}
		}

		if len(garbage) > 0 && !strings.HasSuffix(garbage, "mul(") {
			continue
		}

		nums := strings.Split(data, ",")

		if len(nums) != 2 {
			continue
		}

		num1, err := strconv.Atoi(nums[0])

		if err != nil {
			continue
		}

		num2, err := strconv.Atoi(nums[1])

		if err != nil {
			continue
		}

		result += num1 * num2

	}

	fmt.Println("Part 1 ->", result)
}

func part2(lines []string) {
	reader := bufio.NewReader(strings.NewReader(strings.Join(lines, "")))

	var result int
	enabled := true

	for {
		startBytes, err := reader.ReadBytes(byte('('))

		if errors.Is(err, io.EOF) {
			break
		} else {
			utils.CheckError(err)
		}

		startBytesStr := string(startBytes)

		isInstruction := strings.HasSuffix(startBytesStr, "do(") || strings.HasSuffix(startBytesStr, "don't(")

		if isInstruction {
			b, err := reader.ReadByte()
			utils.CheckError(err)
			if b == ')' {
				enabled = !strings.HasSuffix(startBytesStr, "don't(")
			}
		}

		if !strings.HasSuffix(string(startBytes), "mul(") {
			continue
		}

		endBytes, err := reader.ReadBytes(byte(')'))

		if errors.Is(err, io.EOF) {
			break
		} else {
			utils.CheckError(err)
		}

		endBytes = endBytes[:len(endBytes)-1]

		var data string
		var garbage string

		for _, b := range endBytes {
			if unicode.IsDigit(rune(b)) || b == ',' {
				data += string(b)
			} else {
				data = ""
				garbage += string(b)
			}
		}

		if len(garbage) > 0 && strings.HasSuffix(garbage, "don't(") {
			enabled = false
			continue
		}

		if len(garbage) > 0 && strings.HasSuffix(garbage, "do(") {
			enabled = true
			continue
		}

		if len(garbage) > 0 && !strings.HasSuffix(garbage, "mul(") {
			continue
		}

		if !enabled {
			continue
		}

		nums := strings.Split(data, ",")

		if len(nums) != 2 {
			continue
		}

		num1, err := strconv.Atoi(nums[0])

		if err != nil {
			continue
		}

		num2, err := strconv.Atoi(nums[1])

		if err != nil {
			continue
		}

		result += num1 * num2

	}

	fmt.Println("Part 2 ->", result)
}
