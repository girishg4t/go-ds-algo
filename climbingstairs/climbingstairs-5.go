package climbingstairs

//1
// 1
// 2
// 1 1 , 2
// 3
// 1 1 1, 2 1, 1 2
// 4
// 1 1 1 1, 2 1 1, 1 2 1, 1 1 2, 2 2
// 5
// 1 1 1 1 1, 2 1 1 1, 1 2 1 1, 1 1 2 1, 1 1 1 2
// 6
// 1 1 1 1 1 1, 2 1 1 1 1, 1 2 1 1 1 , 1 1 2 1 1, 1 1 1 2 1, 1 1 1 1 2

/*
 1, 1, 2 , 3, 5
*/
var i = 0

func ClimbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := 1
	b := 1
	c := a + b
	for i := 2; i < n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}
