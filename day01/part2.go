package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	sort.Ints(listA)
	sort.Ints(listB)

	var similarity, ai, bi int

outer:
	for ; ai < len(listA); ai++ {
		a := listA[ai]

		for a >= listB[bi] {
			if a == listB[bi] {
				similarity += a
			}

			bi++

			if bi >= len(listB) {
				break outer
			}
		}
	}

	fmt.Printf("Part 2: %d\n", similarity)
}
