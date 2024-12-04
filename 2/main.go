package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) (reports [][]int) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	datStr := strings.Split(string(dat), "\n")
	reports = make([][]int, len(datStr))

	for i := 0; i < len(datStr); i++ {
		v := strings.Split(datStr[i], " ")
		report := make([]int, len(v))
		for j := 0; j < len(v); j++ {
			report[j], _ = strconv.Atoi(v[j])
		}
		reports[i] = report
	}
	return reports
}

func solvePart1(reports [][]int) int {
	var safeReports int
out:
	for i := 0; i < len(reports); i++ {
		report := reports[i]
		direction := (report[0] - report[1]) > 0
		for j := 1; j < len(report); j++ {
			delta := report[j-1] - report[j]
			if delta == 0 {

				continue out
			}
			newDirection := (delta) > 0
			if newDirection != direction {
				continue out
			}
			if delta < 0 {
				delta *= -1
			}
			if delta > 3 {
				continue out
			}
		}
		safeReports += 1
	}
	return safeReports
}

func checkReport(report []int, toSkip int) (bool, int) {
	// we only want to recheck once (can only remove one element)
	recheckable := true
	if toSkip != -1 {
		var tempReport []int
		for i := 0; i < len(report); i++ {
			if i != toSkip {
				tempReport = append(tempReport, report[i])
			}
		}
		report = tempReport
		recheckable = false
	}
	startingDirection := (report[0] - report[1]) > 0
	for j := 1; j < len(report); j++ {
		result := true
		delta := report[j-1] - report[j]
		currentDirection := (report[j-1] - report[j]) > 0
		if ((delta == 0) || (delta > 3) || (delta < -3)) || (currentDirection != startingDirection) {
			result = false
		}
		if !result {
			if recheckable {
				accum := false
				for k := 0; k < len(report); k++ {
					tempResult, _ := checkReport(report, k)
					accum = (accum || tempResult)
				}
				if !accum {
					return false, j
				}
			} else {
				return false, j
			}
		}
	}
	return true, -1
}

func solvePart2(reports [][]int) int {
	safeReports := 0
	for i := 0; i < len(reports); i++ {
		fmt.Printf("i: %d\n", i)
		result, _ := checkReport(reports[i], -1)
		if result {
			safeReports++
		}
	}
	return safeReports
}

func main() {
	input := os.Args[1]
	reports := readInput(input)
	part1soln := solvePart1(reports)
	fmt.Printf("Part 1 solution is: %d\n", part1soln)
	part2soln := solvePart2(reports)
	fmt.Printf("Part 2 solution is: %d\n", part2soln)
}
