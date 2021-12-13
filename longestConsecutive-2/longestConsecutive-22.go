package longestconsecutive

import (
	"math"
)

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	seq := make(map[int]bool, len(nums))
	for _, n := range nums {
		seq[n] = true
	}
	for _, n := range nums {
		if !seq[n-1] && !seq[n+1] {
			delete(seq, n)
		}
	}
	max := 1
	for val, _ := range seq {
		newMax := 1
		num := val + 1
		for seq[num] {
			newMax++
			num++
		}
		max = int(math.Max(float64(max), float64(newMax)))
	}
	return max
}
