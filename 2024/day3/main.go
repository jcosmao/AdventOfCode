package main

import (
	aoc "aoc/lib"
	"fmt"
	"os"
	"regexp"
)

func main() {
	config := aoc.ParseFlags()
	lines := aoc.ReadFile(config.File)

	switch config.Part {
	case 1:
		processPart1(lines)
	case 2:
		processPart2(lines)
	default:
		os.Exit(1)
	}
}

type FileFormat struct {
	Lines []string
}

func parseLines(lines []string) FileFormat {

	f := FileFormat{}

	for _, line := range lines {
		f.Lines = append(f.Lines, line)
	}

	return f
}

func processPart1(lines []string) {
	file := parseLines(lines)
	fmt.Println(file)

	total := 0
	for _, mul := range getMultiplicationResults(file, true) {
		total += mul
	}

	fmt.Println("part 1: ", int(total))
}

func processPart2(lines []string) {
	file := parseLines(lines)
	fmt.Println(file)

	total := 0
	for _, mul := range getMultiplicationResults(file, false) {
		total += mul
	}

	fmt.Println("part 2: ", int(total))
}

func getMultiplicationResults(file FileFormat, all bool) []int {

	var res []int

	r := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	r_is_mul := regexp.MustCompile(`^mul.*`)

	var mul_enabled bool = true

	for _, line := range file.Lines {
		findall := r.FindAllStringSubmatch(line, -1)
		for _, find := range findall {
			if r_is_mul.MatchString(find[0]) && (mul_enabled || all) {
				a := aoc.StringToInt(find[2])
				b := aoc.StringToInt(find[3])
				res = append(res, a*b)

			} else if find[0] == "don't()" {
				mul_enabled = false
			} else if find[0] == "do()" {
				mul_enabled = true
			}
		}
	}

	return res
}
