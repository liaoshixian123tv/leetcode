package question1833

import (
	"sort"
)

func MaxIceCream(costs []int, coins int) int {
	var count int
	sort.Ints(costs)
	for _, v := range costs {
		if coins >= v {
			coins = coins - v
			count++
		} else {
			break
		}
	}
	return count
}
