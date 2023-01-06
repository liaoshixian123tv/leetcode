package question15

import (
	"sort"
)

func ThreeSumSolution1(nums []int) [][]int {
	var result [][]int
	var left, middle, right int
	sort.Ints(nums)
	for middle = 1; middle < len(nums)-1; middle++ {
		left = 0
		right = len(nums) - 1
		// 倘若middle跟上一個數字重複、則無須重新判斷 只需再判斷middle對應值重複時 只需判斷 middle + middle + right = 0
		if middle > 0 && nums[middle] == nums[middle-1] {
			left = middle - 1
		}
		for left < middle && right > middle {
			// 去掉左邊的重複
			if left > 0 && nums[left] == nums[left-1] {
				left++
				continue
			}
			// 去掉右邊的重複
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			sum := nums[left] + nums[middle] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[left], nums[middle], nums[right]})
				left++
				right--
			} else if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			}

		}

	}
	return result
}

func ThreeSumSolution2(nums []int) [][]int {
	sort.Ints(nums)
	nums = unique(nums)
	var result [][]int
	if len(nums) < 3 {
		return result
	}
	var left, middle, right int
	left, middle, right = 0, 1, 2

	tmpMap := make(map[int]map[int]map[int]bool)
	if tmpMap == nil {
		tmpMap = make(map[int]map[int]map[int]bool)
	}
	for left <= len(nums)-3 {
		if nums[left]+nums[middle]+nums[right] == 0 {
			arr := []int{nums[left], nums[middle], nums[right]}
			// sort.Ints(arr)
			if tmpMap[arr[0]] == nil {
				tmpMap[arr[0]] = make(map[int]map[int]bool)
			}
			if tmpMap[arr[0]][arr[1]] == nil {
				tmpMap[arr[0]][arr[1]] = make(map[int]bool)
			}

			if !tmpMap[arr[0]][arr[1]][arr[2]] {
				tmpMap[arr[0]][arr[1]][arr[2]] = true
				result = append(result, []int{nums[left], nums[middle], nums[right]})
			}

		}
		if right == len(nums)-1 {
			if middle == len(nums)-2 {
				if left == len(nums)-3 {
					return result
				} else {
					left++
				}
				middle = left + 1
			} else {
				middle++
			}
			right = middle + 1
		} else {
			right++
		}
	}

	return result
}

func reorder(num1, num2, num3 int) []int {
	a := []int{num1, num2, num3}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a
}

func unique(arr []int) (result []int) {
	tmpMap := make(map[int]bool)
	for _, v := range arr {
		if !tmpMap[v] {
			result = append(result, v)
			tmpMap[v] = true
		}
	}
	return
}
