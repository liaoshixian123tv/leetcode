package question6

/*
解題思路
主要是對順序的控制
假設numRows是5 並且把他視為行數
則0-> 1 -> 2 -> 3 -> 4 -> 3 -> 2 -> 1 -> 0 -> 1 -> 2.......
所以0至4為正序   4至0為倒序
最後再與input string 的 idx對應可以得到以下
  0-> 1 -> 2 -> 3 -> 4 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11.......
*/

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var tmp string
	tmpMap := make(map[int]string)
	var idx, row int
	var order bool = true
	for idx < len(s) {
		if row == 0 {
			order = true
		} else if row == numRows-1 {
			order = false
		}
		tmpMap[row] += string(s[idx])

		if order {
			row++
		} else {
			row--
		}

		idx++
	}

	for i := 0; i < numRows; i++ {
		tmp += tmpMap[i]
	}
	return tmp
}
