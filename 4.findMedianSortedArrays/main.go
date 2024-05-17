package main

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	findMedianSortedArrays(nums1, nums2)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	nums := make([]int, 0, len1+len2)
	lena := 0
	if len1 > len2 {
		lena = len1
	} else {
		lena = len2
	}
	var offset1, offset2 = 0, 0
	for i := 0; i < lena; {
		for offset1 < len1 || offset2 < len2 {
			if offset1 == len1 {
				nums = append(nums, nums2[offset2])
				offset2++
				i++
				continue
			}
			if offset2 == len2 {
				nums = append(nums, nums1[offset1])
				offset1++
				i++
				continue
			}

			if nums1[offset1] < nums2[offset2] {
				nums = append(nums, nums1[offset1])
				offset1++
				i++
			} else {
				nums = append(nums, nums2[offset2])
				offset2++
				i++
			}
		}
	}
	return sum(nums)
}

func sum(data []int) float64 {
	length := len(data)
	var median float64
	if length%2 == 0 {
		median = float64(data[length/2-1]+data[length/2]) / 2
	} else {
		median = float64(data[length/2])
	}
	return median
}
