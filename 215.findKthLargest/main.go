package main

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4
*/
func main() {
	//nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	nums := []int{4, 3, 2, 5}
	k := findKthLargest(nums, 1)
	println(k)
}



func findKthLargest(nums []int, k int) int {
	TopKSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func Parition(nums []int, start, stop int)int{
	if start >= stop{
		return -1
	}
	pivot := nums[start]
	l, r := start, stop
	for l < r{
		for l < r && nums[r] >= pivot {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] < pivot {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = pivot
	return l
}

func TopKSplit(nums []int, k, start, stop int){
	if start < stop{
		index := Parition(nums, start, stop)

		if index == k {
			return
		} else if index < k {
			TopKSplit(nums, k,index+1, stop )
		} else if index > k {
			TopKSplit(nums, k , start , index - 1)
		}
	}
}

