package checkduplicate

/*
Example 1:

Input: nums = [1,2,3,1]
Output: true

Example 2:

Input: nums = [1,2,3,4]
Output: false

Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Approach :
map[int]int


*/

func CheckDuplicate(nums []int) bool {
	var arr map[int]int = make(map[int]int)
	i := 0
	j := len(nums) - 1
	if len(nums) <= 1 {
		return true
	}
	for i < j {
		a := nums[i]
		b := nums[j]
		if a != b && arr[a] == 0 && arr[b] == 0 {
			arr[a]++
			arr[b]++
			i++
			j--
		} else {
			return true
		}

	}
	return false
}
