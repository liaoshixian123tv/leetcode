package question39

import "sort"

func CombinationSum(candidates []int, target int) [][]int {

	var result [][]int
	var currentPath []int
	sort.Ints(candidates)

	var backtrack func(startIdx int, remainTarget int)
	backtrack = func(currentStartIdx int, currentRemainTarget int) {

		if currentRemainTarget < 0 {
			return
		}

		if currentRemainTarget == 0 {
			tmp := make([]int, len(currentPath))
			copy(tmp, currentPath)
			result = append(result, tmp)
			return
		}

		for i := currentStartIdx; i < len(candidates); i++ {
			if i > currentStartIdx && candidates[i] == candidates[i-1] {
				continue
			}
			currentPath = append(currentPath, candidates[i])

			backtrack(i+1, currentRemainTarget-candidates[i])

			currentPath = currentPath[:len(currentPath)-1]
		}

	}

	backtrack(0, target)

	return result
}
