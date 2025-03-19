package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"math"
	"os"
	"reflect"
	"regexp"
	"runtime/pprof"
	"strings"
)

// Day
type Day struct {
	config  *aoc.Config
	A       int
	B       int
	C       int
	Program []int
	stdout  []string
}

// The adv instruction (opcode 0) performs division. The numerator is the value in the A register. The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of the division operation is truncated to an integer and then written to the A register.

// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand, then stores the result in register B.

// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits), then writes that value to the B register.

// The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero, it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.

// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C, then stores the result in register B. (For legacy reasons, this instruction reads an operand but ignores it.)

// The out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value. (If a program outputs multiple values, they are separated by commas.)

// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the B register. (The numerator is still read from the A register.)

// The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the C register. (The numerator is still read from the A register.)

func combo(operand int, day *Day) int {
	if operand <= 3 {
		return operand
	} else if operand == 4 {
		return day.A
	} else if operand == 5 {
		return day.B
	} else if operand == 6 {
		return day.C
	} else {
		panic("Not valid")
	}
}

func operate(opcode int, operand int, day *Day) int {

	switch opcode {
	case 0:
		combo := combo(operand, day)
		result := float64(day.A) / math.Pow(float64(2), float64(combo))
		day.A = int(result)
	case 1:
		day.B = day.B ^ operand
	case 2:
		combo := combo(operand, day)
		// bitwise AND  with 111 mask
		result := (combo % 8) & 7
		day.B = result
	case 3:
		if day.A != 0 {
			return operand
		}
	case 4:
		// XOR B^C
		day.B = day.B ^ day.C
	case 5:
		combo := combo(operand, day)
		result := combo % 8
		day.stdout = append(day.stdout, aoc.IntToString(result))
	case 6:
		combo := combo(operand, day)
		result := float64(day.A) / math.Pow(float64(2), float64(combo))
		day.B = int(result)
	case 7:
		combo := combo(operand, day)
		result := float64(day.A) / math.Pow(float64(2), float64(combo))
		day.C = int(result)
	}

	return 2
}

func getResult(d *Day) string {

	d.stdout = []string{}

	regA := d.A

	i := 0

	for i < len(d.Program) {
		opcode := d.Program[i]
		operand := d.Program[i+1]

		next := operate(opcode, operand, d)
		if next > len(d.Program) {
			l.Info("END")
			break
		}

		l.Debug("Day", "", d)

		if next != 2 {
			// jump
			l.Debug("Jump", "to", next)
			i = next
		} else {
			i += 2
		}

	}

	if d.config.Part == 1 {
		return strings.Join(d.stdout[:], ",")
	} else if d.config.Part == 2 {
		if reflect.DeepEqual(aoc.IntStringToDigits(strings.Join(d.stdout, "")), d.Program) {
			return aoc.IntToString(regA)
		}
	}

	return ""
}

func parseLines(lines []string, day *Day) {

	rr := regexp.MustCompile(`Register (\w): (\d+)`)
	rp := regexp.MustCompile(`Program: (.*)`)

	for _, line := range lines {
		register := rr.FindStringSubmatch(line)
		if len(register) > 0 {
			if register[1] == "A" {
				day.A = aoc.StringToInt(register[2])
			} else if register[1] == "B" {
				day.B = aoc.StringToInt(register[2])
			} else if register[1] == "C" {
				day.C = aoc.StringToInt(register[2])
			}
		}
		program := rp.FindStringSubmatch(line)
		if len(program) > 0 {
			prog := strings.Split(program[1], ",")
			for _, v := range prog {
				day.Program = append(day.Program, aoc.StringToInt(v))
			}
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
	fmt.Printf("part %d: %s\n", config.Part, res)
}
