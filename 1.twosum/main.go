package main

import "fmt"

func main() {
 fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for index, num := range nums {
		a, ok := numsMap[target-num]
		if ok {
			return []int{a, index}
		} else {
			numsMap[num] = index
		}
	}
	return []int{0,0}
}