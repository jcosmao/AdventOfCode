package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"os"
	"runtime/pprof"
	"slices"
	"strconv"
	"strings"
)

// Day
type Day struct {
	Map           *aoc.Grid
	Summits       []aoc.Position
	// for each start position, list all reachable summit
	StartToSummit map[aoc.Position][]aoc.Position
	// for each start position, save a list of possible tracks (cursor)
	AllTracks     map[aoc.Position][]aoc.Cursor
}

// Height:  Node Interface represent Summit height
type Height struct {
	val int
}

func (i Height) Repr() string {
	return strconv.Itoa(i.val)
}

func (i Height) GetValue() int {
	return i.val
}


func GoHiking(c aoc.Cursor, m *Day) {
	allDirections := []aoc.Direction{aoc.N, aoc.E, aoc.S, aoc.W}
	for _, dir := range allDirections {
		nextPos := c.Grid.GetDirection(c.Pos, dir)
		if nextPos == c.Pos {
			// cannot mv to dir
			continue
		}

		// next mv has +1 value
		if c.Grid.Nodes[*nextPos].GetValue() == c.Grid.Nodes[*c.Pos].GetValue()+1 {
			// duplicate cursor in case there is another possible track from here
			cPrev := c
			c.Move(dir)

			if slices.Contains(m.Summits, *c.Pos) {
				// summit reached
				start := c.GetStart()
				m.AllTracks[*start] = append(m.AllTracks[*start], c)

				// summt has not been visited yet from start position ?
				if !slices.Contains(m.StartToSummit[*start], *c.Pos) {
					m.StartToSummit[*start] = append(m.StartToSummit[*start], *c.Pos)
				}
			}

			GoHiking(c, m)
			// Try next directions with copy of cursor at previous position
			c = cPrev
		}
	}
}

func getResult(d *Day, part int) int {
	total := 0

	if part == 1 {
		for _, v := range d.StartToSummit {
			total += len(v)
		}

	} else if part == 2 {
		for start := range d.AllTracks {
			total += len(d.AllTracks[start])
		}
	}

	return total
}

func parseMap(lines []string) *Day {

	day := &Day{
		AllTracks:     map[aoc.Position][]aoc.Cursor{},
		Summits:       []aoc.Position{},
		StartToSummit: map[aoc.Position][]aoc.Position{},
	}
	day.Map = &aoc.Grid{
		Cols:  0,
		Rows:  len(lines),
		Nodes: make(map[aoc.Position]aoc.T),
	}

	for i, line := range lines {
		splitted := strings.Split(line, "")
		day.Map.Cols = len(splitted)

		for j := 0; j < len(splitted); j++ {
			p := aoc.Position{X: j, Y: i}
			height := aoc.StringToInt(splitted[j])
			day.Map.Nodes[p] = Height{val: height}

			// start
			if height == 0 {
				day.AllTracks[p] = []aoc.Cursor{}
				day.StartToSummit[p] = []aoc.Position{}
			}

			// summit
			if height == 9 {
				day.Summits = append(day.Summits, p)
			}
		}
	}

	return day
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

	lines := aoc.ReadFile(config.File)
	d := parseMap(lines)
	l.Debug("Parsed Map", "map", d)

	for start := range d.AllTracks {
		c := aoc.Cursor{
			Pos:   &start,
			Grid:   d.Map,
			Track: []*aoc.Position{&start},
		}

		l.Debug("Start Hiking", "from", start)
		GoHiking(c, d)
	}

	res := getResult(d, config.Part)
	fmt.Printf("part %d: %d", config.Part, res)
}
