package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"os"
	"runtime/pprof"
	"strings"
)

// Day
type Day struct {
	config aoc.Config
}

func getResult(d *Day) int {
	total := 0

	if d.config.Part == 1 {
	} else if d.config.Part == 2 {
	}

	return total
}

func parseLines(lines []string, day *Day) {

	for i, line := range lines {
		sline := strings.Split(line, "")
	}
}

func main() {
	config := aoc.ParseFlags()

	if config.Profiling {
		_ = os.Remove("cpu.prof")
		cpuf, _ := os.Create("cpu.prof")
		pprof.StartCPUProfile(cpuf)
		defer cpuf.Close()
		defer pprof.StopCPUProfile()
		l.Info("[profiling] go tool pprof cpu.prof")
	}
	defer aoc.Timer("main")()

	lines := aoc.ReadFile(config.File)
	day := &Day{
		config: config,
	}

	parseLines(lines, day)
	l.Debug("Parsed input", "Day", day)

	res := getResult(day)
	fmt.Printf("part %d: %d\n", config.Part, res)
}
