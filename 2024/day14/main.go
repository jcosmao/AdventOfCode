package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"os"
	"regexp"
	"runtime/pprof"
	"time"
)

// Day
type Day struct {
	Robots []*Robot
	Map    *aoc.Grid
}

type Robot struct {
	c      *aoc.Cursor
	iter   int
	vector aoc.Position
}

func getResult(d *Day, part int) int {
	total := 0

	if part == 1 {
		for i := 1; i <= 100; i++ {
			for _, robot := range d.Robots {
				robot.c.MoveVector(robot.vector)
			}

			time.Sleep(time.Second/2)
			aoc.ClearScreen()
			d.Map.Display()
		}

		// split map / 4
		areas := []int{0, 0, 0, 0}

		for _, robot := range d.Robots {
			if robot.c.Pos.X < d.Map.Cols/2 && robot.c.Pos.Y < d.Map.Rows/2 {
				areas[0] += 1
			} else if robot.c.Pos.X > d.Map.Cols/2 && robot.c.Pos.Y < d.Map.Rows/2 {
				areas[1] += 1
			} else if robot.c.Pos.X < d.Map.Cols/2 && robot.c.Pos.Y > d.Map.Rows/2 {
				areas[2] += 1
			} else if robot.c.Pos.X > d.Map.Cols/2 && robot.c.Pos.Y > d.Map.Rows/2 {
				areas[3] += 1
			}
		}

		l.Info("Areas", "", areas)

		total = areas[0] * areas[1] * areas[2] * areas[3]

	} else if part == 2 {

		for i := 1; i <= 10000; i++ {
			for _, robot := range d.Robots {
				robot.c.MoveVector(robot.vector)
			}

			successive := 0
			maxsuccessive := 0
			for j := 0; j < d.Map.Rows; j++ {
				for k := 0; k < d.Map.Cols; k++ {
					pos := aoc.Position{X: k, Y: j}
					_, exists := d.Map.Node[pos]
					if exists && len(d.Map.Node[pos]) > 0 {
						successive += 1
					} else {
						if successive > maxsuccessive {
							maxsuccessive = successive
						}
						successive = 0
					}
				}
				successive = 0
			}
			if maxsuccessive > 10 {
				l.Info("Iter", "", i)
				d.Map.Display()
			}
		}

	}

	return total
}

func parseLines(lines []string) *Day {

	day := &Day{}

	r := regexp.MustCompile(`p=(\d+),(\d+) v=(-*\d+),(-*\d+)`)

	m := aoc.NewGrid(103, 101)
	day.Map = m

	for _, line := range lines {
		f := r.FindStringSubmatch(line)
		l.Debug("Regex", "line", line, "match", f)
		robot := new(Robot)
		if len(f) > 0 {
			px := aoc.StringToInt(f[1])
			py := aoc.StringToInt(f[2])
			vx := aoc.StringToInt(f[3])
			vy := aoc.StringToInt(f[4])

			robot.c = new(aoc.Cursor)
			robot.c.Str = "#"
			robot.c.Grid = m
			robot.c.Pos = new(aoc.Position)
			robot.c.Pos.X = px
			robot.c.Pos.Y = py
			robot.vector = aoc.Position{X: vx, Y: vy}

			m.Node[*robot.c.Pos] = []*aoc.Cursor{robot.c}

			l.Debug("robot", "", robot, "pos", robot.c.Pos.ToString())

			day.Robots = append(day.Robots, robot)
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
	d := parseLines(lines)

	l.Debug("Parsed input", "Day", d)

	res := getResult(d, config.Part)
	fmt.Printf("part %d: %d", config.Part, res)
}
