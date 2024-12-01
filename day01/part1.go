package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func Part1() {
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

	var distance int
	for i := 0; i < len(listA); i++ {
		distance += int(math.Abs(float64(listA[i] - listB[i])))
	}

	fmt.Printf("Part 1: %d\n", distance)
}
