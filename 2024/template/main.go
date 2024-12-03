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
	Lines [][]int
}

func parseLines(lines []string) FileFormat {

	f := FileFormat{}

	for _, line := range lines {
		splitted := strings.Fields(line)

		vals := []int{}
		for i := 0; i <= 1; i++ {
			r, err := strconv.Atoi(splitted[i])
			if err != nil {
				panic(err)
			}
			vals = append(vals, r)
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
