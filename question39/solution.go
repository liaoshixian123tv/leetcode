package question39

// func CombinationSum(candidates []int, target int) [][]int {

// 	var result [][]int
// 	var currentPath []int

// 	var backtrack func(startIdx int, remainTarget int)
// 	backtrack = func(currentStartIdx int, currentRemainTarget int) {

// 		if currentRemainTarget < 0 {
// 			return
// 		}

// 		if currentRemainTarget == 0 {
// 			tmp := make([]int, len(currentPath))
// 			copy(tmp, currentPath)
// 			result = append(result, tmp)
// 			return
// 		}

// 		for i := currentStartIdx; i < len(candidates); i++ {
// 			currentPath = append(currentPath, candidates[i])

// 			backtrack(i, currentRemainTarget-candidates[i])

// 			currentPath = currentPath[:len(currentPath)-1]
// 		}

// 	}

// 	backtrack(0, target)

// 	return result
// }

// [1,2,3,4,5,6,7] 6
func CombinationSum(candidates []int, target int) [][]int {

	var result [][]int
	var currentPath []int

	var backtrack func(startIdx int, remainTarget int)
	backtrack = func(currentStartIdx, currentRemainTarget int) {

		if currentRemainTarget < 0 {
			return
		}

		if currentRemainTarget == 0 {
			tmp := make([]int, len(currentPath))
			copy(tmp, currentPath)
			result = append(result, tmp)
		}

		for i := currentStartIdx; i < len(candidates); i++ {
			currentPath = append(currentPath, candidates[i])
			backtrack(i, currentRemainTarget-candidates[i])
			currentPath = currentPath[:len(currentPath)-1]
		}

	}

	backtrack(0, target)

	return result
}
