package question9

import "strconv"

/*
解題思路:
把數字反轉後看看原始數字使否相同

*/

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmpStr := ""
	tmpX := strconv.Itoa(x)

	for i := len([]byte(tmpX)) - 1; i >= 0; i-- {
		tmpStr += string([]byte(tmpX)[i])
	}
	if tmpStr == tmpX {
		return true
	}
	return false

}
