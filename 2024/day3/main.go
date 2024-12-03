package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

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

	var total int = 0
	for _, mul := range getMultiplicationResults(file, true) {
		total += mul
	}

	fmt.Println("part 1: ", int(total))
}

func processPart2(lines []string) {
	file := parseLines(lines)
	fmt.Println(file)

	var total int = 0
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
		// fmt.Printf("%#v\n", findall)
		for _, find := range findall {
			if r_is_mul.MatchString(find[0]) && (mul_enabled || all) {
				a, err := strconv.Atoi(find[2])
				if err != nil {
					panic(err)
				}

				b, err := strconv.Atoi(find[3])
				if err != nil {
					panic(err)
				}

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
