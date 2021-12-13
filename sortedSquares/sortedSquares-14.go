package sortedsquares

import "math"

/*
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
[16,1]
[0,9,100]

*/
// [0,1, 9,16,100]
func SortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	i := 0
	j := n - 1

	for n = n - 1; n > -1; n-- {
		if math.Abs(float64(nums[i])) < math.Abs(float64(nums[j])) {
			result[n] = nums[j] * nums[j]
			j--
		} else {
			result[n] = nums[i] * nums[i]
			i++
		}
	}
	return result
}
