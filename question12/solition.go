package question12

/*
	I can be placed before V (5) and X (10) to make 4 and 9.
	X can be placed before L (50) and C (100) to make 40 and 90.
	C can be placed before D (500) and M (1000) to make 400 and 900.

	1 <= num <= 3999

	解題思路:
	主要先定義好特殊數字並做好對應關係
	再從數組裡從最大的數字不斷相減，倘若相減後的input小於最大數字、再從數組裡挑下一個做為最大數字以此類推
	直到input = 0 代表相減完成

*/

func IntToRoman(num int) string {
	var intArr []int = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var romanArr []string = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var idx int
	var result string
	for num > 0 {
		if num >= intArr[idx] {
			num = num - intArr[idx]
			result += romanArr[idx]
		} else {
			idx++
		}
	}

	return result
}
