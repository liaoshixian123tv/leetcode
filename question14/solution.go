package question14

/*
解題思路:
先取input裡最短長度的字串作為max length
再從0開始作為每個字串的idx
倘若每個字串的字母一樣 則新增到result
反之則直接return
*/

func LongestCommonPrefix(strs []string) string {
	var minVal, idx int
	var result []byte
	if len(strs) == 0 {
		return ""
	}

	minVal = len(strs[0])
	for _, v := range strs {
		minVal = min(minVal, len(v))
	}

	for idx < minVal {
		tmpS := strs[0][idx]
		var isSame bool = true
		for _, s := range strs {
			if tmpS != s[idx] {
				isSame = false
			}
		}
		if isSame {
			result = append(result, tmpS)
		} else {
			return string(result)
		}
		idx++
	}

	return string(result)
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
