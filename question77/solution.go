package question77

/*
Given two integers n and k, return all possible combinations of k numbers
chosen from the range [1, n]. You may return the answer in any order.

Example 1:

Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.
Example 2:

Input: n = 1, k = 1
Output: [[1]]
Explanation: There is 1 choose 1 = 1 total combination.


Constraints:

1 <= n <= 20
1 <= k <= n

*/

// combine n: max number, k: elements
func Combine(n int, k int) [][]int {
	var result [][]int
	var currentPath []int
	var start = 1

	var trackback func(startIdx int)
	trackback = func(currentStartIdx int) {

		if len(currentPath) > k {
			return
		}

		if len(currentPath) == k {
			tmp := make([]int, k)
			copy(tmp, currentPath)
			result = append(result, tmp)
			return
		}

		for i := currentStartIdx; i <= n-(k-len(currentPath))+1; i++ {
			currentPath = append(currentPath, i)
			trackback(i + 1)
			currentPath = currentPath[:len(currentPath)-1]
		}

	}

	trackback(start)

	return result
}
