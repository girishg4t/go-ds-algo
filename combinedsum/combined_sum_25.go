package combinedsum

import "sort"

//2,3,5
func CombinationSum(candidates []int, target int) [][]int {
	store := make(map[int]bool)
	result := [][]int{}
	for _, n := range candidates {
		store[n] = true
	}

	for _, n := range candidates {
		sum := target
		subset := []int{}
		for {
			if sum < 0 {
				break
			}
			sum = sum - n
			if sum == 0 {
				subset = append(subset, n)
				result = append(result, subset)
				break
			} else {
				if store[sum-n] {
					subset = append(subset, n)
					subset = append(subset, sum-n)
					result = append(result, subset)
				}
				if store[sum] && !checkDuplicate(result, n, sum, subset) {
					subset = append(subset, n)
					subset = append(subset, sum)
					result = append(result, subset)
					return result
				}
			}

			subset = append(subset, n)
		}
	}
	return result
}

func checkDuplicate(result [][]int, n, sum int, subset []int) bool {
	subset = append(subset, n)
	subset = append(subset, sum)

	for _, elm := range result {
		sort.Slice(elm, func(i, j int) bool {
			return i > j
		})
		sort.Slice(subset, func(i, j int) bool {
			return i > j
		})
		if len(subset) == len(elm) {
			for i, _ := range elm {
				if elm[i] == subset[i] {
					continue
				} else {
					return false
				}
			}
			return true
		}
	}
	return false
}
