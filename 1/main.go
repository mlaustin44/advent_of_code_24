package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(filename string) (list1, list2 []int) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	datStr := strings.Split(string(dat), "\n")
	list1 = make([]int, len(datStr))
	list2 = make([]int, len(datStr))

	for i := 0; i < len(datStr); i++ {
		v := strings.Split(datStr[i], "   ")
		list1[i], _ = strconv.Atoi(v[0])
		list2[i], _ = strconv.Atoi(v[1])
	}
	return list1, list2
}

func solvePart1(list1, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)

	var sum int

	for i := 0; i < len(list1); i++ {
		dist := list1[i] - list2[i]
		if dist < 0 {
			dist *= -1
		}
		sum += dist
	}
	return sum
}

func solvePart2(list1, list2 []int) int {
	frequency := make(map[int]int)
	for i := 0; i < len(list2); i++ {
		frequency[list2[i]] += 1
	}
	var similSum int
	for i := 0; i < len(list1); i++ {
		v := list1[i]
		f := frequency[v]
		sim := v * f
		similSum += sim
	}
	return similSum
}

func main() {
	input := os.Args[1]
	list1, list2 := readInput(input)
	part1soln := solvePart1(list1, list2)
	fmt.Printf("Part 1 solution is: %d\n", part1soln)
	part2soln := solvePart2(list1, list2)
	fmt.Printf("Part 2 solution is: %d\n", part2soln)
}
