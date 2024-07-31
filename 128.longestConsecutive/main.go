package main

func main() {
	longestConsecutive([]int{0, 3, 2, 4})
}

func longestConsecutive(nums []int) int {
	numsMap := map[int]bool{}
	if len(nums) <= 0 {
		return 0
	}
	for _, num := range nums {
		numsMap[num] = true
	}
	ret := 1
	for _, num := range nums {
		if !numsMap[num-1] {
			i := 0
			curNum := num
			for numsMap[curNum] {
				i++
				curNum++
			}

			if i > ret {
				ret = i
			}
		}
	}

	return ret
}
