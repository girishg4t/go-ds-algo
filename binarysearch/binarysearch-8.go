package binarysearch

func Search(nums []int, target int) int {
	m := len(nums) / 2
	i := m
	for i <= m {
		if nums[m] == target {
			return i
		}
		if nums[m] < target {
			i = m + 1
			m = len(nums)
		}
		if nums[m] > target {
			m = m - 1
		}
		m = m / 2
	}

	return -1
}
