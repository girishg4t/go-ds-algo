package numbersdisappeared

import (
	"fmt"
	"math"
)

func FindDisappearedNumbers(nums []int) []int {
	var uniqueNums map[int]bool = make(map[int]bool)
	for _, n := range nums {
		uniqueNums[n] = true
	}

	result := []int{}

	for i := 1; i <= len(nums); i++ {
		if !uniqueNums[i] {
			result = append(result, i)
		}
	}
	return result
}

func FindDisappearedNumbers_2(nums []int) []int {

	for _, n := range nums {
		num := int(math.Abs(float64(n))) - 1
		if nums[num] > 0 {
			nums[num] = -nums[num]
		}
	}

	result := []int{}
	fmt.Println(nums)
	for i, _ := range nums {
		if nums[i] > 0 {
			result = append(result, i)
		}
	}
	return result
}
