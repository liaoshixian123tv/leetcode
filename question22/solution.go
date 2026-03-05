package question22

func generateParenthesis(n int) []string {
	var result []string

	var backtrack func(left int, right int, current string)
	backtrack = func(left int, right int, current string) {

		// 1. 終止條件：括號都用完了
		if left == 0 && right == 0 {
			result = append(result, current)
			return
		}

		// 2. 分支一：只要還有左括號，就可以加左括號
		if left > 0 {
			backtrack(left-1, right, current+"(")
		}

		// 3. 分支二：只要剩餘的左括號比右括號少（代表前面有配對的左括號），就可以加右括號
		if left < right {
			backtrack(left, right-1, current+")")
		}
	}

	backtrack(n, n, "")
	return result
}
