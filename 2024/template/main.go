package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"os"
	"runtime/pprof"
	"strings"
)

// FileFormat
type FileFormat struct {
	Lines [][]int
}

func processPart1(lines []string) {
	file := parseLines(lines)
	l.Debug("", "file", file)

	var total int = 0
	fmt.Println("part 1: ", int(total))
}

func processPart2(lines []string) {
	file := parseLines(lines)
	l.Debug("", "file", file)

	var total int = 0
	fmt.Println("part 2: ", int(total))
}

func parseLines(lines []string) FileFormat {

	f := FileFormat{}

	for _, line := range lines {
		splitted := strings.Fields(line)

		vals := []int{}
		for i := 0; i < len(splitted); i++ {
			vals = append(vals, aoc.StringToInt(splitted[i]))
		}
		f.Lines = append(f.Lines, vals)
	}

	return f
}

func main() {
	config := aoc.ParseFlags()
	lines := aoc.ReadFile(config.File)

	if config.Profiling {
		_ = os.Remove("cpu.prof")
		cpuf, _ := os.Create("cpu.prof")
		pprof.StartCPUProfile(cpuf)
		defer cpuf.Close()
		defer pprof.StopCPUProfile()
		l.Info("[profiling] go tool pprof cpu.prof")
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
