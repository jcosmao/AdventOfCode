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
	config      *aoc.Config
	designsMap  map[string]int
	designs     []string
	patternList []string
}

func canBuildDesign(orig string, design string, day *Day) bool {
	if design == "" {
		l.Debug("Found", "design", orig)
		day.designsMap[orig] += 1
		l.Debug("all", "", day.designsMap)
		return true
	}

	for _, pattern := range day.patternList {
		l.Debug("Check pattern", "pattern", pattern, "design", design)
		if strings.HasPrefix(design, pattern) {
			subdesign, _ := strings.CutPrefix(design, pattern)
			if ok := canBuildDesign(orig, subdesign, day); ok {
				l.Debug("Check subdesign", "orig", orig, "pattern", pattern, "design", subdesign)
				// return true
			}
		}
	}

	return false
}

func getResult(day *Day) int {
	total := 0
	total2 := 0
	day.designsMap = make(map[string]int)
	// for each design, check if can be built from a list of patterns
	for _, design := range day.designs {
		l.Debug("Check", "design", design)
		day.designsMap[design] = 0
		canBuildDesign(design, design, day)
		if day.designsMap[design] > 0 {
			total += 1
			total2 += day.designsMap[design]
		}
	}

	if day.config.Part == 1 {
		return total
	} else if day.config.Part == 2 {
		return total2
	}

	return total
}

func parseLines(lines []string, day *Day) {

	day.designs = []string{}
	day.patternList = []string{}

	for i, line := range lines {
		if i == 0 {
			sline := strings.Split(line, ", ")
			day.patternList = sline

		} else {
			if line == "" {
				continue
			}
			day.designs = append(day.designs, strings.TrimSpace(line))
		}
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
