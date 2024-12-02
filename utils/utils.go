package utils

import (
	"flag"
	"os"
	"strconv"
	"strings"
)

func Lines() []string {
	fileName := flag.String("f", "input.txt", "-f <filename>")
	flag.Parse()

	data, err := os.ReadFile(*fileName)
	trimmedData := strings.TrimSpace(string(data))
	CheckError(err)

	lines := strings.Split(trimmedData, "\n")

	return Map(lines, func(line string, index int) string {
		return strings.TrimSpace(line)
	})
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Map[I interface{}, O interface{}](arr []I, fn func(I, int) O) []O {
	result := make([]O, len(arr))

	for i, v := range arr {
		result[i] = fn(v, i)
	}

	return result
}

func Filter[I interface{}](arr []I, fn func(I, int) bool) []I {
	result := make([]I, 0)

	for i, v := range arr {
		if fn(v, i) {
			result = append(result, v)
		}
	}

	return result
}

func Reduce[I interface{}, O interface{}](arr []I, fn func(O, I, int) O, initial O) O {
	result := initial

	for idx, v := range arr {
		result = fn(result, v, idx)
	}

	return result
}

func Every[I interface{}](arr []I, fn func(I) bool) bool {
	for _, v := range arr {
		if !fn(v) {
			return false
		}
	}

	return true
}

func MustAtoi(a string) int {
	res, err := strconv.Atoi(a)
	CheckError(err)
	return res
}
