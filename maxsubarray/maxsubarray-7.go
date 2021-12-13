package maxsubarray

import (
	"math"
)

// [-2,1,-3,4,-1,2,1,-5,4]
func MaxSubArray(nums []int) int {
	max := nums[0]
	if len(nums) == 1 {
		max = nums[0]
	}
	sum := nums[0]
	for _, n := range nums[1:] {
		sum = sum + n
		sum = int(math.Max(float64(sum), float64(n)))
		if sum < 0 {
			sum = n
		}
		max = int(math.Max(float64(max), float64(sum)))
	}
	return max
}
