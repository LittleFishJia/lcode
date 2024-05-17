package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring1("abcba"))
}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans

}

func lengthOfLongestSubstring1(s string) int {
	if s == "" {
		return 0
	}
	var maxLen = 1
	var m = make(map[byte]int)
	var offset = -1
	n := len(s)
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for offset+1 < n && m[s[offset+1]] == 0 {
			if m[s[offset+1]] == 0 {
				m[s[offset+1]]++
				offset++
			}
		}
		maxLen = max(maxLen, offset-i+1)
	}
	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
