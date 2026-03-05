package question45

// [2,3,0,1,4]
func Jump(nums []int) int {
	var jumpsCount int = 0
	var currentEnd int = 0
	var farthest int = 0

	for i := 0; i < len(nums)-1; i++ {

		if farthest < i+nums[i] {
			farthest = i + nums[i]
		}

		if i == currentEnd {
			jumpsCount++
			currentEnd = farthest
		}
	}

	return jumpsCount
}
