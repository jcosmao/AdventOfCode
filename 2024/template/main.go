package main

import (
	"aoc"
	"fmt"
	"os"
	"strings"
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

// FileFormat
type FileFormat struct {
	Lines [][]int
}

func parseLines(lines []string) FileFormat {

	f := FileFormat{}

	for _, line := range lines {
		splitted := strings.Fields(line)

		vals := []int{}
		for i := 0; i <= 1; i++ {
			vals = append(vals, aoc.StringToInt(splitted[i]))
		}
		f.Lines = append(f.Lines, vals)
	}

	return f
}

func processPart1(lines []string) {
	file := parseLines(lines)
	fmt.Println(file)

	var total int = 0
	fmt.Println("part 1: ", int(total))
}

func processPart2(lines []string) {
	file := parseLines(lines)
	fmt.Println(file)

	var total int = 0
	fmt.Println("part 2: ", int(total))
}
