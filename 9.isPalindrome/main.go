package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	var arr []int
	for n > 0 {
		b := n % 10
		arr = append(arr, b)
		n = n / 10
	}
	nLen := len(arr)

	for i, j := 0, nLen-1; i <= j; {

		if arr[i] == arr[j] {
			i++
			j--
		} else {
			return false
		}

	}

	return true
}
