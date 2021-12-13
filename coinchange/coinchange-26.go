package coinchange

import (
	"fmt"
	"math"
	"sort"
)

// 1,2,5
func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Slice(coins, func(i, j int) bool {
		return i > j
	})
	sort.Ints(coins)
	fmt.Println(coins)
	total := 0
	store := make(map[int]bool)
	for _, n := range coins {
		store[n] = true
		total = total + n
	}
	prevCoins := -1
	if total == amount {
		prevCoins = len(coins)
	}
	return recurese(coins, amount, prevCoins, store)
}

func recurese(coins []int, amount, prevCoins int, store map[int]bool) int {
	total := amount
	numberOfCoins := 0
	index := len(coins) - 1
	for {
		if index < 0 {
			numberOfCoins = -1
			break
		}
		total = total - coins[index]
		numberOfCoins++
		if total == 0 {
			break
		}
		if total > 0 && store[total] {
			numberOfCoins++
			break
		}
		if total < 0 {
			total = total + coins[index]
			numberOfCoins--
			index--
		}
	}

	if index > -1 {
		if prevCoins == -1 {
			prevCoins = numberOfCoins
		} else {
			prevCoins = int(math.Min(float64(prevCoins), float64(numberOfCoins)))
		}
		return recurese(coins[:index], amount, prevCoins, store)
	}
	return prevCoins
}
