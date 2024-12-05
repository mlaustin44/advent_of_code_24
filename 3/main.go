package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
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
		if in[i:i+3] == "mul" {
			secondParen, comma := -1, -1
			validFound := false
			for j := (i + 4); j < len(in); j++ {
				if string(in[j]) == "," {
					comma = j
				} else if (string(in[j]) == ")") && (comma != -1) {
					secondParen = j
					validFound = true
					break
				} else if (!unicode.IsNumber(rune(in[j]))) || in[j:j+3] == "mul" {
					break
				}

			}
			if validFound {
				num1, _ := strconv.Atoi(in[i+4 : comma])
				num2, _ := strconv.Atoi(in[comma+1 : secondParen])
				sum += num1 * num2
			}
		}
	}
	return sum
}

func solvePart2(in string) int {
	sum := 0
	enabled := true
	for i := 0; i < (len(in) - 2); i++ {
		if (in[i:i+3] == "mul") && enabled {
			secondParen, comma := -1, -1
			validFound := false
			for j := (i + 4); j < len(in); j++ {
				if string(in[j]) == "," {
					comma = j
				} else if (string(in[j]) == ")") && (comma != -1) {
					secondParen = j
					validFound = true
					break
				} else if (!unicode.IsNumber(rune(in[j]))) || in[j:j+3] == "mul" {
					break
				}

			}
			if validFound {
				num1, _ := strconv.Atoi(in[i+4 : comma])
				num2, _ := strconv.Atoi(in[comma+1 : secondParen])
				sum += num1 * num2
			}
		} else if (i < (len(in) - 3)) && (in[i:i+4] == "do()") {
			enabled = true
		} else if (i < (len(in) - 6)) && (in[i:i+7] == "don't()") {
			enabled = false
		}
	}
	return sum
}

func main() {
	input := os.Args[1]
	in := readInput(input)
	part1soln := solvePart1(in)
	fmt.Printf("Part 1 solution is: %d\n", part1soln)
	part2soln := solvePart2(in)
	fmt.Printf("Part 2 solution is: %d\n", part2soln)
}
