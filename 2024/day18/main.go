package main

import (
	aoc "aoc/lib"
	"container/heap"
	"fmt"
	l "log/slog"
	"os"
	"runtime/pprof"
	"strings"
)

type Day struct {
	config *aoc.Config
	grid   *aoc.Grid
	start  aoc.Position
	end    aoc.Position
	costs  map[aoc.Position]map[aoc.Direction]int
	max    int
	a      int
	b      int
	all    []aoc.Position
}

type State struct {
	position  aoc.Position
	cost      int
	direction aoc.Direction
	index     int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func findShortestPath(day *Day) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &State{position: day.start, cost: 0, direction: aoc.E})

	day.costs = make(map[aoc.Position]map[aoc.Direction]int)
	day.costs[day.start] = make(map[aoc.Direction]int)
	day.costs[day.start][0] = 0

	directions := []aoc.Direction{aoc.N, aoc.E, aoc.S, aoc.W}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*State)
		l.Debug("Process", "node", current)

		if current.position == day.end {
			l.Debug("costs", "", day.costs)
			return current.cost
		}

		for _, direction := range directions {
			if direction == current.direction.Opposite() {
				continue
			}

			newPos := day.grid.GetDirection(&current.position, direction)
			if newPos != nil && newPos != &current.position {
				if c, exists := day.grid.Node[*newPos]; exists {
					// wall
					l.Debug("node", "", c[0].Str)
					continue
				}

				// cost
				newCost := current.cost + 1

				if _, exists := day.costs[*newPos]; !exists {
					day.costs[*newPos] = make(map[aoc.Direction]int)
				}

				if _, exists := day.costs[*newPos][direction]; !exists || newCost < day.costs[*newPos][direction] {
					day.costs[*newPos][direction] = newCost
					heap.Push(pq, &State{position: *newPos, cost: newCost, direction: direction})
				}
			}
		}
	}

	return -1 // Si aucun chemin n'est trouvé
}

func getResult(day *Day) int {
	return findShortestPath(day)
}

func parseLines(lines []string, day *Day) {

	day.all = []aoc.Position{}
	day.grid = &aoc.Grid{
		Rows: 70,
		Cols: 70,
		Node: make(map[aoc.Position][]*aoc.Cursor),
	}

	if day.config.File == "test.txt" {
		day.grid.Rows = 6
		day.grid.Cols = 6
	}

	for i := 0; i < day.max; i++ {
		pos := strings.Split(lines[i], ",")
		x := aoc.StringToInt(pos[0])
		y := aoc.StringToInt(pos[1])

		position := aoc.Position{X: x, Y: y}
		l.Debug("position", "", position)
		day.grid.Node[position] = []*aoc.Cursor{}
		node := new(aoc.Cursor)
		node.Str = "#"
		day.grid.Node[position] = append(day.grid.Node[position], node)

		day.all = append(day.all, position)
	}

	day.grid.Cols += 1
	day.grid.Rows += 1

	day.start = aoc.Position{X: 0, Y: 0}
	day.end = aoc.Position{X: day.grid.Cols - 1, Y: day.grid.Rows - 1}
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
		a:      0,
		b:      len(lines),
	}

	res := -1

	if day.config.Part == 1 {

		// break after X position
		if day.config.File == "test.txt" {
			day.max = 12
		} else {
			day.max = 1024
		}
		parseLines(lines, day)
		res = getResult(day)
		fmt.Printf("part %d: %d\n", config.Part, res)

	} else {
		// dichotomie
		for {
			if day.a+2 == day.b {
				day.max = len(lines)
				parseLines(lines, day)

				l.Debug("Found", "a", day.a, "b", day.b, "First fail", day.all[day.a])
				break
			}

			test := day.a + ((day.b - day.a) / 2)

			day.max = test
			parseLines(lines, day)
			res = getResult(day)

			if res == -1 {
				day.b = test + 1
			} else {
				day.a = test
			}
		}

		fmt.Printf("part %d: %+v\n", config.Part, day.all[day.a])
	}
}

