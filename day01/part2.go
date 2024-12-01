package main

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	listA := make([]int, 1000)
	listB := make([]int, 1000)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		var a, b int
		fmt.Sscanf(line, "%d   %d", &a, &b)

		listA = append(listA, a)
		listB = append(listB, b)
	}

	counts := make(map[int]int)

	for i := 0; i < len(listA); i++ {
		counts[listA[i]] = 0
	}

	for i := 0; i < len(listB); i++ {
		if _, ok := counts[listB[i]]; ok {
			counts[listB[i]]++
		}
	}

	similarity := 0
	for a, b := range counts {
		similarity += a * b
	}

	fmt.Printf("Part 2: %d\n", similarity)
}
