package main

import "fmt"

/*

33. 搜索旋转排序数组
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
														在下标 3 处经旋转后可能变为 [2,4,5,6,7,0,2]
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
*/

func main() {
	nums := []int{1,3,5}
	a := search(nums, 5)
	fmt.Println(a)
}

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	gap := nums[0]
	maxIndex := findMax(nums)
	fmt.Println(maxIndex)
	if target > nums[maxIndex] {
		return -1
	}
	if target == nums[maxIndex] {
		return maxIndex
	}

	if target > gap {
		for i:=0 ;i< maxIndex; i++ {
			if nums[i] == target {
				return i
			}
		}
	} else if target< gap{
		for i:=len(nums) - 1 ;i> maxIndex; i-- {
			if nums[i] == target {
				return i
			}
		}
	}else  {

		return 0
	}
	return -1
}

func findMax(nums []int)  int {
	mid := len(nums) / 2
	maxIndex := 0
	for i := 0; i <= mid;  i++ {
		if nums[i] > nums[maxIndex] {//左边都是递增
			maxIndex = i
		}
	}

	for j := mid+1; j <= len(nums) - 1;  j++ {
		if nums[j] > nums[maxIndex] {//左边都是递增
			maxIndex =j
		}
	}
	return maxIndex
}


// 153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	for left+1 < right { // 开区间不为空
		mid := left + (right-left)/2
		if nums[mid] < nums[len(nums)-1] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

// 有序数组中找 target 的下标
func lowerBound(nums []int, left, right, target int) int {
	for left+1 < right { // 开区间不为空
		// 循环不变量：
		// nums[left] < target
		// nums[right] >= target
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid // 范围缩小到 (mid, right)
		} else {
			right = mid // 范围缩小到 (left, mid)
		}
	}
	if nums[right] != target {
		return -1
	}
	return right
}

