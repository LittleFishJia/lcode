package main

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组
是数组中的一个连续部分。
示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/
func main()  {
	nums := []int {-2,1,-3,4,-1,2,1,-5,4}
	maxSubArray(nums)
}

func maxSubArray(nums []int) ([]int, int) {
	if len(nums) == 0 {
		return []int{}, 0
	}
	maxSoFar := nums[0]
	maxEndingHere := nums[0]
	l, r := 0,0
	for i:= 1 ; i< len(nums); i++{
		if nums[i] > maxEndingHere + nums[i] {
			maxEndingHere = nums[i]
			l = i
		} else {
			maxEndingHere += nums[i]
		}
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
			r = i
		}
	}
	return nums[l:r+1],maxSoFar
}
