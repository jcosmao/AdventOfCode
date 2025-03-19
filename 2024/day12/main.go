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
	Map          *aoc.Grid
	GardenPlants map[string][]*Garden
}

// Height:  Node Interface represent Summit height
type Node struct {
	position aoc.Position
	letter   string
	// nb of fence needed.    nb fence + nb same letter adj = 4
	val     int
	visited bool
	edges   []*Node
}

type Garden struct {
	letter    string
	members   []*Node
	dirChange int
}

func (i Node) Repr() string {
	return i.letter
}

func (i Node) GetValue() int {
	return i.val
}

func (i Node) CountSides() int {
	// todo
	return 0
}


func getResult(d *Day, part int) int {
	total := 0

	for _, Inode := range d.Map.Nodes {
		n := Inode.(*Node)
		garden := Garden{letter: n.Repr(), members: []*Node{}, dirChange:0}
		l.Debug("Create New garden", "garden", garden)
		Gardening(n, &garden, d)
		l.Debug("Finished Gardening", "garden", garden, "len", len(garden.members))

		if len(garden.members) > 0 {
			d.GardenPlants[garden.letter] = append(d.GardenPlants[garden.letter], &garden)
		}
	}

	if part == 1 {
		for _, v := range d.GardenPlants {
			for _, garden := range v {
				fence := 0
				for i := 0; i < len(garden.members); i++ {
					n := garden.members[i]
					fence += n.val
					l.Debug("Node", "", n)
				}
				total += len(garden.members) * fence
			}
		}
	} else if part == 2 {
		for _, v := range d.GardenPlants {
			for _, garden := range v {
				fence := 0
				for i := 0; i < len(garden.members); i++ {
					fence += garden.members[i].CountSides()
				}
				total += len(garden.members) * fence
			}
		}

	}

	return total
}

func Gardening(node *Node, garden *Garden, d *Day) {

	if node.visited {
		return
	}

	node.visited = true

	l.Debug("Add node to garden", "node", node)
	garden.members = append(garden.members, node)
	l.Debug("Garden members", "garden", garden)

	allPos := d.Map.GetAllDirections(&node.position)

	for _, neighbor := range allPos {
		edge := d.Map.Nodes[*neighbor].(*Node)

		if edge.Repr() == node.Repr() {
			node.val -= 1
			node.edges = append(node.edges, edge)
			l.Debug("Checking neighbor", "edge", edge.position.ToString(), "letter", edge.letter)
			Gardening(edge, garden, d)
		}
	}

}

func parseMap(lines []string) *Day {

	day := &Day{
		GardenPlants: make(map[string][]*Garden),
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
			node := &Node{letter: splitted[j], position: p, val: 4}
			day.Map.Nodes[p] = node
			if _, exists := day.GardenPlants[node.letter]; !exists {
				day.GardenPlants[node.letter] = []*Garden{}
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
	for letter, v := range d.GardenPlants {
		l.Debug("Gardens", "letter", letter, "nb", len(v))
	}

	res := getResult(d, config.Part)
	fmt.Printf("part %d: %d", config.Part, res)
}
