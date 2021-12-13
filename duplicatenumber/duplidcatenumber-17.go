package duplicatenumber

import "fmt"

// [1,2,3,4,2]
// 10

// [1,2,3,]
// n *(n+1)
// 3 * 4 = 12/2
// 4*5= 20/2 = 10

// [2,2,2,2,2]
// [1,2,3,4,4,4]
//
func FindDuplicate(nums []int) int {
	var slow = nums[0]
	var fast = nums[nums[0]]

	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
		fmt.Println(slow, fast)
	}

	slow = 0
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
		fmt.Println(slow, fast)
	}
	return slow
}
