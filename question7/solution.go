package question7

import (
	"math"
)

/*
 -231 <= x <= 231 - 1
解題思路:

主要取數字的最後一位並且取得後下次再*10再加上下次的最後一位

注意反轉後的數字也要在範圍內
-231 <= result <= 231 - 1
*/
func Reverse(x int) int {

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	var isNegative bool
	if x < 0 {
		isNegative = true
		x *= -1
	}

	var result int
	for x > 0 {
		num := x % 10
		result = result*10 + num
		x = x / 10
	}
	if isNegative {
		result *= -1
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}
