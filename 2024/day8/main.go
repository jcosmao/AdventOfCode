package main

import (
	"fmt"
	l "log/slog"
	"os"
	"runtime/pprof"
	"slices"
	"strings"

	aoc "aoc/lib"
)

const (
	ANTI string = "#"
)

type Position struct {
	x int
	y int
}

type Antenna struct {
	signal string
}

type Map struct {
	rows      int
	cols      int
	nodes     map[Position]*Antenna
	antennas  map[string][]Position
	conflicts []Position
}

func GetSymmetric(a Position, b Position) Position {

	new := Position{}

	opposite := func(c int, d int) int {
		res := 0
		if c < d {
			res = d + (d - c)
		} else if c == d {
			res = c
		} else {
			res = d - (c - d)
		}
		return res
	}

	new.x = opposite(a.x, b.x)
	new.y = opposite(a.y, b.y)

	l.Debug("Symmetric", "a", a, "b", b, "new", new)
	return new
}

func (m *Map) Display() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			pos := Position{x: j, y: i}
			if m.nodes[pos] == nil {
				fmt.Print(".")
			} else {
				fmt.Print(m.nodes[pos].signal)
			}
		}
		fmt.Println("")
	}
}

func (m *Map) GetResonancePoints(a Position, b Position) []Position {
	res := []Position{}

	for {
		c := GetSymmetric(a, b)
		if !m.OutOfBound(c) {
			res = append(res, c)
			a = b
			b = c
		} else {
			break
		}
	}
	return res
}

func (m *Map) OutOfBound(p Position) bool {
	if p.x < 0 || p.x >= m.cols {
		return true
	}

	if p.y < 0 || p.y >= m.rows {
		return true
	}

	return false
}

func (m *Map) AddAntinode(p Position) {

	if m.OutOfBound(p) {
		return
	}

	if m.nodes[p] == nil {
		// add antinode
		antinode := Antenna{signal: ANTI}
		m.antennas[antinode.signal] = append(m.antennas[antinode.signal], p)
		m.nodes[p] = &antinode
	} else {
		l.Error("Conflict", "", p, "found", m.nodes[p].signal)
		if m.nodes[p].signal != ANTI {
			if !slices.Contains(m.conflicts, p) {
				m.conflicts = append(m.conflicts, p)
			}
		}
	}
}

func processPart1(lines []string) {
	m := parseLines(lines)
	// l.Debug("", "file", m)
	m.conflicts = []Position{}

	for nodePos, nodeAnt := range m.nodes {
		l.Debug("node", "pos", nodePos, "val", nodeAnt)
		if nodeAnt != nil {
			// it's an antenna - Need to find antinode for each
			if nodeAnt.signal == ANTI {
				continue
			}
			for _, pos := range m.antennas[nodeAnt.signal] {
				if pos == nodePos {
					continue
				}
				antinodePos := GetSymmetric(nodePos, pos)
				m.AddAntinode(antinodePos)
			}
		}
	}

	m.Display()

	total := len(m.antennas[ANTI]) + len(m.conflicts)

	fmt.Println("part 1: ", int(total))
}

func processPart2(lines []string) {
	m := parseLines(lines)

	m.conflicts = []Position{}

	for nodePos, nodeAnt := range m.nodes {
		l.Debug("node", "pos", nodePos, "val", nodeAnt)
		if nodeAnt != nil {
			// it's an antenna - Need to find antinode for each
			if nodeAnt.signal == ANTI {
				continue
			}
			for _, pos := range m.antennas[nodeAnt.signal] {
				if pos == nodePos {
					continue
				}
				antinodes := m.GetResonancePoints(nodePos, pos)
				for _, a := range antinodes {
					m.AddAntinode(a)
				}
			}
		}
	}

	m.Display()

	total := 0
	for _, v := range m.antennas {
		total += len(v)
	}

	// total := len(m.antennas[ANTI]) + len(m.conflicts)
	fmt.Println("part 2: ", int(total))
}

func parseLines(lines []string) Map {

	m := Map{}
	m.rows = len(lines)
	m.nodes = map[Position]*Antenna{}
	m.antennas = map[string][]Position{}

	for i, line := range lines {
		s := strings.Split(line, "")
		if m.cols == 0 {
			m.cols = len(s)
		}

		for j := 0; j < len(s); j++ {
			pos := Position{x: j, y: i}
			if s[j] != "." {
				ant := Antenna{signal: s[j]}
				m.nodes[pos] = &ant
				m.antennas[ant.signal] = append(m.antennas[ant.signal], pos)
			} else {
				m.nodes[pos] = nil
			}
		}

	}

	return m
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
