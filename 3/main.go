package main

import (
	"fmt"
	"os"
	"strconv"
)

func readInput(filename string) string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(dat)
}

func solvePart1(in string) int {
	sum := 0
	for i := 0; i < (len(in) - 2); i++ {
		fmt.Printf("i: %d\n", i)
		if in[i:i+3] == "mul" {
			firstParen, secondParen, comma := -1, -1, -1
			validFound := false
			for j := (i + 3); j < len(in); j++ {
				fmt.Printf("\tj: %d, in[j]: %s, firstParen: %d, secondParen: %d, comma: %d\n", j, string(in[j]), firstParen, secondParen, comma)
				if (string(in[j]) == "(") && (comma == -1) {
					firstParen = j
				} else if (string(in[j]) == ",") && (firstParen != -1) {
					comma = j
				} else if (string(in[j]) == ")") && ((firstParen != -1) && (comma != -1)) {
					secondParen = j
					validFound = true
					break
				} else if in[j:j+3] == "mul" {
					break
				}
			}
			if validFound {
				num1, _ := strconv.Atoi(in[firstParen+1 : comma])
				num2, _ := strconv.Atoi(in[comma+1 : secondParen])
				fmt.Printf("found multiple instruction for %d and %d", num1, num2)
				sum += num1 * num2
			}
		}
	}
	return sum
}

func main() {
	input := os.Args[1]
	in := readInput(input)
	part1soln := solvePart1(in)
	fmt.Printf("Part 1 solution is: %d\n", part1soln)
	// part2soln := solvePart2(in)
	// fmt.Printf("Part 2 solution is: %d\n", part2soln)
}
