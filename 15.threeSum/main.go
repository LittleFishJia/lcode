package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, -2, -1}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for a := 0; a < n; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		b := a + 1
		c := n - 1
		for b != c && b < c {
			if b > a+1 && nums[b] == nums[b-1] {
				b++
				continue
			}
			if nums[c]+nums[a]+nums[b] > 0 {
				c--
				continue
			}
			if nums[c]+nums[a]+nums[b] == 0 {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
			b++

		}

	}
	return ans
}
