package findsubset

import "sort"

// [[],[9],[0],[0,9],[3],[3,9],[0,3],[0,3,9]
func Subsets(nums []int) [][]int {
	result := [][]int{{}}
	sort.Slice(nums, func(i, j int) bool {
		return i > j
	})
	backtrack(result, []int{}, nums, 0)
	return result
}

func backtrack(result [][]int, tempList []int, nums []int, start int) {
	result = append(result, tempList)
	for i := start; i < len(nums); i++ {
		tempList = append(tempList, nums[i])
		backtrack(result, tempList, nums, i+1)
		tempList = tempList[:len(tempList)-1]
	}
}
