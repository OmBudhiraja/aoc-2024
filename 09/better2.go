package main

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/ombudhiraja/aoc-2024/utils"
)

type MinHeap []int

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func betterPart2(line string) {
	var result int

	// create a min heap for for spaces 0 to 10
	indexHeaps := make([]MinHeap, 10)

	disk := make([]string, 0)

	for i := range line {
		if i%2 != 0 {
		}

	}

	for i, ch := range line {
		num := utils.MustAtoi(string(ch))

		if i%2 != 0 {
			if num == 0 {
				continue
			}
			heap.Push(&indexHeaps[num], len(disk))
			disk = append(disk, repeat(".", num)...)
		} else {
			id := fmt.Sprintf("%d", i/2)
			disk = append(disk, repeat(id, num)...)
		}

	}

	i := len(disk) - 1

	for i >= 0 {
		if disk[i] == "." {
			i--
			continue
		}

		fileId := disk[i]
		fileCount := 0

		for i >= 0 && disk[i] == fileId {
			fileCount++
			i--
		}

		smallestIdx := math.MaxInt
		bestWidth := -1

		for width := fileCount; width < len(indexHeaps); width++ {
			if len(indexHeaps[width]) == 0 {
				continue
			}

			if indexHeaps[width][0] < smallestIdx && indexHeaps[width][0] < i {
				bestWidth = width
				smallestIdx = indexHeaps[width][0]
			}
		}

		if smallestIdx == math.MaxInt {
			continue
		}

		// remove from the used heap and push the remaining spaces to appropriate heap
		heap.Pop(&indexHeaps[bestWidth])
		heap.Push(&indexHeaps[bestWidth-fileCount], smallestIdx+fileCount)

		for j := 0; j < fileCount; j++ {
			disk[smallestIdx+j] = fileId
			disk[i+1+j] = "."
		}

	}

	for i, block := range disk {
		if block == "." || block == "_" {
			continue
		}

		result += utils.MustAtoi(block) * i
	}

	fmt.Println("part 2 better ->", result)
}
