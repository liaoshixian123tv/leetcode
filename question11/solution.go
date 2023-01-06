package question11

/*
	解題思路:
	主要分別為數組的頭、尾一個指針
	並且每次取兩指針最小值作為高
	兩指針距離作為寬
	面積 = 高 * 寬 並且記錄下來
	每當面積計算完成後，將兩只針作為高的指針往另一個指針移動，直到兩針相遇為止
	最後取最大的面積即可

	n == height.length
	2 <= n <= 105
	0 <= height[i] <= 104
*/

func MaxArea(height []int) int {
	start, end, maxArea := 0, len(height)-1, 0
	for start < end {
		width := end - start
		tmpHeight := 0
		if height[start] < height[end] {
			tmpHeight = height[start]
			start++
		} else {
			tmpHeight = height[end]
			end--
		}
		if width*tmpHeight > maxArea {
			maxArea = width * tmpHeight
		}
	}
	return maxArea
}
