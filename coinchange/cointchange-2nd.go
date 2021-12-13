package coinchange

import (
	"math"
	"sort"
)

// 1 3 4 5 => 7
/*
	[ 0 8 8 8 8 8 8 8]
	[ 0 1 8 8 8 8 8 8]
	[ 0 1 2 1 8 8 8 8]
	[ 0 1 2 1 1 8 8 8]
	[ 0 1 2 1 1 1 8 8]
	[ 0 1 2 1 1 1 2 8]
	[ 0 1 2 1 1 1 2 3]
*/

func CoinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	result := make([]int, amount+1)
	for i, _ := range result {
		result[i] = amount + 1
	}
	sort.Ints(coins)
	result[0] = 0
	for i, _ := range result {
		for _, coin := range coins {
			amt := i - coin
			if amt < 0 {
				break
			}
			if amt == 0 {
				result[i] = 1
			} else {
				if result[amt] > 0 {
					result[i] = int(math.Min(float64(result[i]), float64(result[amt]+1)))
				}
			}
		}
	}
	if result[amount] != amount+1 {
		return result[amount]
	}
	return -1
}
