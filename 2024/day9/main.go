package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"os"
	"runtime/pprof"
)

// FileFormat
type FileFormat struct {
	diskmap []int
}

type File struct {
	fd     int
	len    int
	blocks []*Block
}

type Block struct {
	index int
	data  *File
}

type Disk struct {
	allocationMap map[int]*Block
	used          []int
	free          []int
}

func processPart1(lines []string) {
	file := parseLines(lines)
	l.Debug("", "file", file)

	disk := Disk{
		allocationMap: map[int]*Block{},
		used:          []int{},
		free:          []int{},
	}

	fd := 0
	seek := 0

	for i, value := range file.diskmap {
		l.Debug("Seek", "seek", seek)
		if aoc.IsPair(i) {
			l.Debug("Found file", "len", value)
			f := File{fd: fd, len: value, blocks: []*Block{}}
			for j := seek; j < seek+f.len; j++ {
				l.Debug("Set alloc disk space", "index", j, "block", f)
				b := Block{index: j, data: &f}
				f.blocks = append(f.blocks, &b)
				disk.allocationMap[j] = &b
				disk.used = append(disk.used, j)
			}
			fd += 1
			seek += f.len
		} else {
			l.Debug("Found free space", "len", value)
			for j := seek; j < seek+value; j++ {
				l.Debug("Set free disk space", "append", j)
				disk.free = append(disk.free, j)
			}
			seek += value
		}
	}

	l.Debug("Disk", "", disk)

	for i := len(disk.used) - 1; i >= 0; i-- {
		l.Debug("Move", "block", disk.used[i], "target", disk.free[0])
		first_free := disk.free[0]
		last_used := disk.used[i]

		if first_free >= last_used {
			break
		}

		disk.allocationMap[first_free] = disk.allocationMap[last_used]
		delete(disk.allocationMap, last_used)
		disk.free = disk.free[1:]
	}

	l.Debug("Disk", "", disk)

	total := 0
	for i := 0; i < len(disk.allocationMap); i++ {
		total += i * disk.allocationMap[i].data.fd
	}

	fmt.Println("part 1: ", int(total))
}

func processPart2(lines []string) {
	file := parseLines(lines)
	l.Debug("", "file", file)

	disk := Disk{
		allocationMap: map[int]*Block{},
		used:          []int{},
		free:          []int{},
	}

	fd := 0
	seek := 0

	for i, value := range file.diskmap {
		l.Debug("Seek", "seek", seek)
		if aoc.IsPair(i) {
			l.Debug("Found file", "len", value)
			f := File{fd: fd, len: value, blocks: []*Block{}}
			for j := seek; j < seek+f.len; j++ {
				l.Debug("Set alloc disk space", "index", j, "block", f)
				b := Block{index: j, data: &f}
				f.blocks = append(f.blocks, &b)
				disk.allocationMap[j] = &b
			}
			disk.used = append(disk.used, seek)
			fd += 1
			seek += f.len
		} else {
			l.Debug("Found free space", "len", value)
			for j := seek; j < seek+value; j++ {
				l.Debug("Set free disk space", "append", j)
				disk.free = append(disk.free, j)
			}

			seek += value
		}
	}

	l.Debug("Disk", "", disk)

	maxindex := 0
	for k := range disk.allocationMap {
    	if k > maxindex {
			maxindex = k
		}
	}

	for i := len(disk.used) - 1; i >= 0; i-- {
		l.Debug("Move", "file", disk.used[i])
		file := disk.allocationMap[disk.used[i]].data
		// find available place on disk
		for k := 0; k <= maxindex; k++ {
		nextindex:

			l.Debug("check space", "", k)

			for i := 0; i < file.len; i++ {
				if _, exists := disk.allocationMap[k+i]; exists {
					k += 1
					goto nextindex
				}
			}

			l.Debug("Found space", "index", k)

			if file.blocks[0].index < k {
				break
			}

			// move file
			for i := 0; i < file.len; i++ {
				block := Block{index: k + i, data: file}
				disk.allocationMap[k+i] = &block
				delete(disk.allocationMap, file.blocks[i].index)
				file.blocks[i] = &block
			}
			break
		}
	}

	l.Debug("Disk", "", disk)

	total := 0
	for i := 0; i <= maxindex; i++ {
		if _, exists := disk.allocationMap[i]; exists {
			total += i * disk.allocationMap[i].data.fd
		}
	}

	fmt.Println("part 2: ", int(total))
}

func parseLines(lines []string) FileFormat {

	f := FileFormat{}
	f.diskmap = aoc.IntStringToDigits(lines[0])
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
