package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

// Day
type Day struct {
	config     *aoc.Config
	directions []aoc.Direction
	grid       *aoc.Grid
	robot      *aoc.Cursor
	moved      []*aoc.Cursor
}

func MovePush(node *aoc.Cursor, dir aoc.Direction, pair bool, day *Day) bool {
	posNext := node.GetDirection(dir)
	nodeNext := node.Grid.GetFirstNode(*posNext)

	if nodeNext != nil && (nodeNext.Str == "O" || nodeNext.Str == "[" || nodeNext.Str == "]") {
		if ok := MovePush(nodeNext, dir, false, day); !ok {
			return false
		}
		nodeNext = node.Grid.GetFirstNode(*posNext)
	}

	if nodeNext != nil && nodeNext.Str == "#" {
		return false
	}

	if nodeNext == nil {

		if pair == false && (node.Str == "[" || node.Str == "]") && (dir == aoc.N || dir == aoc.S) {

			l.Debug("Move", "cursor", node.Str, "from", node.Pos.ToString(), "to", dir)

			nodePair := new(aoc.Cursor)
			if node.Str == "[" {
				pairPos := node.GetDirection(aoc.E)
				nodePair = node.Grid.GetFirstNode(*pairPos)
			} else {
				pairPos := node.GetDirection(aoc.W)
				nodePair = node.Grid.GetFirstNode(*pairPos)
			}

			l.Debug("Move PAIR", "cursor", node.Str, "pair", nodePair.Str, "from", node.Pos.ToString(), "and", nodePair.Pos.ToString(), "to", dir)

			if ok := MovePush(nodePair, dir, true, day); !ok {
				return false
			}
		}

		l.Debug("Move", "cursor", node.Str, "from", node.Pos.ToString(), "to", dir, "Pair", pair)

		if ok := node.Move(dir, false); !ok {
			return false
		}

		day.moved = append(day.moved, node)

		return true
	}

	return false
}

func getResult(day *Day) int {
	total := 0

	l.Info("START")
	day.grid.Display()

	for _, move := range day.directions {
		l.Debug("Robot", "pos", day.robot.Pos.ToString(), "move", move)

		if ok := MovePush(day.robot, move, false, day); !ok {
			l.Debug("Revert", "", day.moved)
			for _, node := range day.moved {
				node.Move(move.Opposite())
			}
		}

		day.moved = []*aoc.Cursor{}

		if day.config.Display {
			day.grid.Display()
			time.Sleep(time.Second / 25)
			aoc.ClearScreen()
		}
	}

	l.Info("RESULT")
	day.grid.Display()

	for k := range day.grid.Node {
		node := day.grid.GetFirstNode(k)
		if node != nil && (node.Str == "[" || node.Str == "O") {
			total += 100*node.Pos.Y + node.Pos.X
		}
	}

	return total
}

func parseLines(lines []string, day *Day) {

	m := &aoc.Grid{
		Rows: 0,
		Node: make(map[aoc.Position][]*aoc.Cursor),
	}

	for i, line := range lines {
		sline := strings.Split(line, "")
		l.Debug("line", "splitted", sline, "i", i)
		if i == 0 {
			m.Cols = len(sline) * day.config.Part
		}

		if len(sline) > 0 && sline[0] == "#" {

			// duplicate
			if day.config.Part == 2 {
				newsline := []string{}
				for _, char := range sline {
					if char == "." || char == "#" {
						newsline = append(newsline, char)
						newsline = append(newsline, char)
					} else if char == "O" {
						newsline = append(newsline, "[")
						newsline = append(newsline, "]")
					} else if char == "@" {
						newsline = append(newsline, char)
						newsline = append(newsline, ".")
					}
				}
				l.Debug("New line", "", newsline)
				sline = newsline
			}

			m.Rows += 1
			for j, char := range sline {
				pos := aoc.Position{X: j, Y: i}
				m.Node[pos] = []*aoc.Cursor{}
				if char != "." {
					node := &aoc.Cursor{Str: char, Pos: &pos, Grid: m}
					m.Node[pos] = append(m.Node[pos], node)
					if char == "@" {
						day.robot = node
					}
				}
			}
		} else {
			if len(sline) == 0 {
				continue
			}

			for _, move := range sline {
				if move == "v" {
					day.directions = append(day.directions, aoc.S)
				} else if move == ">" {
					day.directions = append(day.directions, aoc.E)
				} else if move == "^" {
					day.directions = append(day.directions, aoc.N)
				} else if move == "<" {
					day.directions = append(day.directions, aoc.W)
				}
			}
		}

		day.grid = m
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

	day := &Day{
		config: config,
	}

	lines := aoc.ReadFile(config.File)
	parseLines(lines, day)

	res := getResult(day)

	fmt.Printf("part %d: %d\n", config.Part, res)
}
