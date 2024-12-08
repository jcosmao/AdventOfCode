package lib

import (
	"math"
	"strconv"
	"strings"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// "1234" -> [1, 2, 3 ,4]
func IntStringToDigits(s string) []int {
	fields := strings.Split(s, "")
	output := make([]int, len(fields))
	for i := 0; i < len(fields); i++ {
		output[i] = StringToInt(fields[i])
	}
	return output
}

// 1234 -> [1, 2, 3, 4]
func IntToDigits(n int64, seq ...[]int64) []int64 {
	sequence := []int64{}
	if seq != nil {
		sequence = seq[0]
	}
    if n != 0 {
        i := n % 10
        // sequence = append(sequence, i) // reverse order output
        sequence = append([]int64{i}, sequence...)
        return IntToDigits(n/10, sequence)
    }
    return sequence
}

// ConcatInts concatenates a list of integers into a single integer.
// 22 345 6 -> 223456
func ConcatInts(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for _, num := range nums[1:] {
		// Calculate the number of digits in num
		numDigits := int(math.Log10(float64(num))) + 1

		// Shift result to the left by the number of digits in num
		result = result * int(math.Pow(10, float64(numDigits))) + num
	}

	return result
}
