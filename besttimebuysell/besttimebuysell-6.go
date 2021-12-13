package besttimebuysell

//  [7,1,5,3,6,4]

// [3,2,6,5,0,3]
// 7, 1, 5, 3, 6, 4
func MaxProfit(prices []int) int {
	result := 0
	bought := prices[0]

	for i := 1; i < len(prices); i++ {
		if bought > prices[i] {
			bought = prices[i]
		} else {
			dif := prices[i] - bought
			if result < dif {
				result = dif
			}
		}
	}

	return result
}
