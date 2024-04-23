package main

import "fmt"

func main() {
	a := "cbbd"
	fmt.Println(longestPalindrome(a))

}
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
