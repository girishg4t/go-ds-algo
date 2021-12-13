package singlenumber

// 2 2 1

func SingleNumber(nums []int) int {
	var mapItem map[int]int = make(map[int]int)
	for _, n := range nums {
		mapItem[n]++
		if mapItem[n] == 2 {
			delete(mapItem, n)
		}
	}
	var result int
	for key, _ := range mapItem {
		result = key
		break
	}
	return result
}
