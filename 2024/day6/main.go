package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"os"
	"runtime/pprof"
	"slices"
	"strings"
)

type Map struct {
	rows    int
	cols    int
	start   Position
	allowed map[Position]bool
}

type Position struct {
	x int
	y int
}

type Guard struct {
	pos       Position
	visited   map[Position][]string
	direction string
	exit      Position
	loop      bool
}

func parseMap(lines []string) Map {

	m := Map{}
	m.rows = len(lines)
	m.allowed = map[Position]bool{}

	for i, line := range lines {
		s := strings.Split(line, "")
		if m.cols == 0 {
			m.cols = len(s)
		}

		for j, v := range s {
			pos := Position{x: j, y: i}

			if v == "." {
				m.allowed[pos] = true
			} else if v == "#" {
				m.allowed[pos] = false
			} else if v == "^" {
				m.start = pos
				m.allowed[pos] = true
			} else {
				l.Warn("Unknown char", "char", v)
			}
		}
	}

	return m
}

func (g *Guard) Walk(m Map) bool {

	newPos := Position{}
	if g.direction == "up" {
		newPos.x = g.pos.x
		newPos.y = g.pos.y - 1
	} else if g.direction == "down" {
		newPos.x = g.pos.x
		newPos.y = g.pos.y + 1
	} else if g.direction == "right" {
		newPos.x = g.pos.x + 1
		newPos.y = g.pos.y
	} else if g.direction == "left" {
		newPos.x = g.pos.x - 1
		newPos.y = g.pos.y
	}

	l.Debug("Walk", "pos", g.pos, "direction", g.direction, "target", newPos)

	val, haskey := m.allowed[newPos]
	if haskey {
		if val {
			g.pos = newPos
			l.Debug("New position", "pos", g.pos)
		} else {
			g.direction = getNextDirection(g.direction)
			l.Debug("Change direction", "direction", g.direction)
			return true
		}
	} else {
		g.exit = newPos
		l.Debug("Found Exit", "exit", g.exit)
		return false
	}

	positions, haskey := g.visited[g.pos]
	if !haskey {
		g.visited[g.pos] = append(g.visited[g.pos], g.direction)
	} else if slices.Contains(positions, g.direction) {
		g.loop = true
	}

	return true
}

func getNextDirection(d string) string {
	dirMap := map[string]string{
		"up":    "right",
		"right": "down",
		"down":  "left",
		"left":  "up",
	}
	return dirMap[d]
}

func processPart1(lines []string) {
	m := parseMap(lines)

	g := &Guard{
		pos:       m.start,
		visited:   map[Position][]string{},
		direction: "up",
	}

	g.visited[m.start] = []string{g.direction}

	for {
		if !g.Walk(m) {
			break
		}
	}

	fmt.Println("part 1: ", len(g.visited))
}

func processPart2(lines []string) {
	m := parseMap(lines)
	total := 0

	g := &Guard{
		pos:       m.start,
		visited:   map[Position][]string{},
		direction: "up",
	}

	// resolve 1 time first to get all visited place
	for {
		if !g.Walk(m) {
			break
		}
	}

	for k := range g.visited {
		// Same player play again
		g.pos = m.start
		g.direction = "up"
		g.visited = map[Position][]string{}
		g.loop = false

		l.Debug("Try and add caillou", "caillou", k)
		m.allowed[k] = false

		for {
			if !g.Walk(m) {
				// exit found, no need to check for loop
				break
			}

			if g.loop {
				l.Debug("Loop detected", "pos", g.pos, "direction", g.direction)
				total += 1
				break
			}
		}

		// remove caillou
		m.allowed[k] = true
	}

	fmt.Println("part 2: ", total)
}

func main() {
	config := aoc.ParseFlags()
	lines := aoc.ReadFile(config.File)

	if config.Profiling {
		f, err := os.Create("cpu.prof")
		if err != nil {
			panic(err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
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
