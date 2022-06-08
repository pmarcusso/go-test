package calc

import (
	"sort"
)

func OrderList(nums []int) []int {
	sort.Ints(nums)
	return nums
}
