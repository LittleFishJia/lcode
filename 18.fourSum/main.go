package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{-494, -487, -471, -470, -465, -462, -447, -445, -441, -432, -429, -422, -406, -398, -397, -364, -344, -333, -328, -307, -302, -293, -291, -279, -269, -269, -268, -254, -198, -181, -134, -127, -115, -112, -96, -94, -89, -58, -58, -58, -44, -2, -1, 43, 89, 92, 100, 101, 106, 106, 110, 116, 143, 156, 168, 173, 192, 231, 248, 256, 281, 316, 321, 327, 346, 352, 353, 355, 358, 365, 371, 410, 413, 414, 447, 473, 473, 475, 476, 481, 491, 498}, 8511))
}

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n <= 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	var result [][]int
	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			start := j + 1
			end := n - 1
			for start < end {
				if start > j+1 && nums[start-1] == nums[start] {
					start++
					continue
				}

				sum := nums[i] + nums[j] + nums[end] + nums[start]
				fmt.Println(i, j, start, end, sum, target)
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[end], nums[start]})
					start++
				}
				if sum < target {
					start++
				}
				if sum > target {
					end--
				}
			}
		}
	}
	return result
}
