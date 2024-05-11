package main

import (
	"fmt"
)

func main()  {
	fmt.Println(romanToInt("VVV"))
}
var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}
