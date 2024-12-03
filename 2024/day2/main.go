package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc/cli"
)

func main() {
	config := cli.ParseFlags()

	lines, err := cli.ReadFile(config.File)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch config.Part {
	case 1:
		processPart1(lines)
	case 2:
		processPart2(lines)
	default:
		os.Exit(1)
	}
}

// FileFormat
type FileFormat struct {
	Reports [][]int
}

func parseLines(lines []string) FileFormat {

	f := FileFormat{}

	for _, line := range lines {
		splitted := strings.Fields(line)

		report := []int{}
		for _, v := range splitted {
			r, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			report = append(report, r)
		}
		f.Reports = append(f.Reports, report)
	}

	return f
}

func processPart1(lines []string) {
	file := parseLines(lines)
	var total int = 0
	for _, r := range file.Reports {
		if ok, _ := isValidReport(r); ok {
			fmt.Println(r)
			total += 1
		}
	}
	fmt.Println("part 1: ", int(total))
}

// isValidReport
func isValidReport(report []int) (bool, int) {
	var prev int
	var status bool = true

	var order string = "inc"
	if report[0] == report[1] {
		return false, 1
	} else if report[0] < report[1] {
		order = "inc"
	} else {
		order = "dec"
	}

	for i, v := range report {
		if i == 0 {
			prev = v
			continue
		}

		// fmt.Println("Check ", prev, " with ", v, " order == ", order)
		if !isNextValid(prev, v, order) {
			return false, i
		}
		prev = v
	}
	return status, -1
}

func isNextValid(current int, next int, order string) bool {
	if order == "inc" && next > current && next <= (current+3) {
		return true
	} else if order == "dec" && next < current && next >= (current-3) {
		return true
	} else {
		return false
	}
}

func processPart2(lines []string) {
	file := parseLines(lines)
	var total int = 0
	for _, r := range file.Reports {
		if ok, _ := isValidReport(r); ok {
			fmt.Println(r)
			total += 1
		} else {
			for i := range r {
				test_report := RemoveIndex(r, i)
				if ok, _ := isValidReport(test_report); ok {
					fmt.Println(r, " is valid with 1 error at index: ", i, " -> ", test_report)
					total += 1
					break
				}
			}
		}
	}
	fmt.Println("part 2: ", int(total))
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
