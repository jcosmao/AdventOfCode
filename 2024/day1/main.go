package main

import (
	aoc "aoc/lib"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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
	Col1 []int
	Col2 []int
}

func parseLines(lines []string) FileFormat {

	f := FileFormat{}

	for _, line := range lines {
		splitted := strings.Fields(line)

		val := []int{}
		for i := 0; i <= 1; i++ {
			r, err := strconv.Atoi(splitted[i])
			if err != nil {
				panic(err)
			}
			val = append(val, r)
		}

		f.Col1 = append(f.Col1, val[0])
		f.Col2 = append(f.Col2, val[1])
	}

	return f
}

func processPart1(lines []string) {
	file := parseLines(lines)

	sort.Sort(sort.IntSlice(file.Col1))
	sort.Sort(sort.IntSlice(file.Col2))

	var total float64 = 0
	var distance float64

	for i := 0; i < len(file.Col1); i++ {
		distance = math.Abs(float64(file.Col1[i]) - float64(file.Col2[i]))
		total += distance
	}

	fmt.Println("part 1: ", int(total))
}

func processPart2(lines []string) {
	file := parseLines(lines)

	sort.Sort(sort.IntSlice(file.Col1))
	sort.Sort(sort.IntSlice(file.Col2))

	occurrencesMap := countOccurrences(file.Col2)

	var total int = 0
	var multi int

	for i := 0; i < len(file.Col1); i++ {

		n := file.Col1[i]
		if val, ok := occurrencesMap[n]; ok {
			multi = n * val
			total += multi
		}
	}

	fmt.Println("part 2: ", total)
}

func countOccurrences(numbers []int) map[int]int {
	occurrences := make(map[int]int)

	for _, number := range numbers {
		occurrences[number]++
	}

	return occurrences
}
