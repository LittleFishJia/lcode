package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	a, b := 0, len(height)-1
	maxArea := 0
	for a < b {
		maxArea = maxNum(minNum(height[a], height[b])*(b-a), maxArea)
		if height[a] <= height[b] {
			a++
		} else {
			b--
		}
	}
	return maxArea
}

func minNum(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxNum(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
