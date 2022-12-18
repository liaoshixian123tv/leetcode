package question3

/*
Given a string s, find the length of the longest
substring without repeating characters.

給一字串並回傳最長的不重複的字母字串
例如： pwwkew 則是 wke 或 kew   所以答案是3
*/
// LengthOfLongestSubstring
func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	var result, left, right int
	tmpMap := make(map[byte]bool)
	for left < len(s) && right < len(s) {
		rightLetter := s[right]
		leftLetter := s[left]
		if !tmpMap[rightLetter] {
			tmpMap[rightLetter] = true
			right++
		} else {
			if tmpMap[leftLetter] {
				tmpMap[leftLetter] = false
			}
			left++
		}
		if (right - left) > result {
			result = right - left
		}
	}
	return result
}
