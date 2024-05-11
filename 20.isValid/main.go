package main

import "fmt"

func main()  {
	fmt.Println(isValid("(("))
}



func isValid(s string)bool  {
	c := len(s)
	if c % 2 == 1 || c == 0 {
		return false
	}
	tmp := map[string]string{
		")":"(",
		"}":"{",
		"]":"[",
	}
	var stacks []string
	for _, a := range s {
		s , ok := tmp[string(a)]
		fmt.Println(stacks, ok,string(a))
		if ok {
			n := len(stacks)
			if n==0 || n-1<0||stacks[n-1] != s {
				return  false
			}
			stacks = stacks[:n-1]
		}else {
			stacks = append(stacks, string(a))
		}
	}
	if len(stacks) > 0 {
		return  false
	}
	return true
}
