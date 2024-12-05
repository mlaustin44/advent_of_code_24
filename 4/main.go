package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) []string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	datStr := string(dat)
	spltStr := strings.Split(datStr, "\n")

	return spltStr
}

func solvePart1(in []string) int {
	// structure to help debug, leaving it in
	debug := make([][]string, len(in))
	for i := 0; i < len(in); i++ {
		debug[i] = make([]string, len(in))
		for j := 0; j < len(in[0]); j++ {
			debug[i][j] = "."
		}
	}

	found := 0
	// rows
	for i := 0; i < len(in); i++ {
		for j := 0; j <= (len(in[0]) - 4); j++ {
			if (in[i][j:j+4] == "XMAS") || (in[i][j:j+4] == "SAMX") {
				found += 1
				debug[i][j] = "X"
				debug[i][j+1] = "X"
				debug[i][j+2] = "X"
				debug[i][j+3] = "X"
			}
		}
	}
	// columns
	for i := 0; i <= (len(in[0]) - 4); i++ {
		for j := 0; j < len(in); j++ {
			subStr := string(in[i][j]) + string(in[i+1][j]) + string(in[i+2][j]) + string(in[i+3][j])
			if (subStr == "XMAS") || (subStr == "SAMX") {
				found += 1
				debug[i][j] = "X"
				debug[i+1][j] = "X"
				debug[i+2][j] = "X"
				debug[i+3][j] = "X"
			}
		}
	}
	// diagonals
	for i := 0; i <= (len(in) - 4); i++ {
		for j := 0; j <= (len(in[0]) - 4); j++ {
			subStr1 := string(in[i][j]) + string(in[i+1][j+1]) + string(in[i+2][j+2]) + string(in[i+3][j+3])

			rJ := len(in[0]) - j - 1
			subStr2 := string(in[i][rJ]) + string(in[i+1][rJ-1]) + string(in[i+2][rJ-2]) + string(in[i+3][rJ-3])

			if (subStr1 == "XMAS") || (subStr1 == "SAMX") {
				found += 1
				debug[i][j] = "X"
				debug[i+1][j+1] = "X"
				debug[i+2][j+2] = "X"
				debug[i+3][j+3] = "X"
			}
			if (subStr2 == "XMAS") || (subStr2 == "SAMX") {
				found += 1
				debug[i][rJ] = "X"
				debug[i+1][rJ-1] = "X"
				debug[i+2][rJ-2] = "X"
				debug[i+3][rJ-3] = "X"
			}
		}
	}
	return found
}

func solvePart2(in []string) int {
	found := 0
	for i := 0; i < (len(in) - 2); i++ {
		for j := 0; j < (len(in[0]) - 2); j++ {
			subStr1 := string(in[i][j]) + string(in[i+1][j+1]) + string(in[i+2][j+2])
			subStr2 := string(in[i][j+2]) + string(in[i+1][j+1]) + string(in[i+2][j])
			if ((subStr1 == "MAS") || (subStr1 == "SAM")) && ((subStr2 == "MAS") || (subStr2 == "SAM")) {
				fmt.Printf("i: %d, j: %d, sub1: %s, sub2: %s\n", i, j, subStr1, subStr2)
				found += 1
			}
		}
	}
	return found
}

func main() {
	input := os.Args[1]
	in := readInput(input)
	part1soln := solvePart1(in)
	fmt.Printf("Part 1 solution is: %d\n", part1soln)
	part2soln := solvePart2(in)
	fmt.Printf("Part 2 solution is: %d\n", part2soln)
}
