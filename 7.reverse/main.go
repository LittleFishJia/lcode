package main

import "fmt"

func main(){
	fmt.Println(reverse(1563847412))
}

func reverse(x int) int {
	if x == 0 {return x}
	sum := 0
	for {
		a , b := x/10, x%10
		if a == 0 {
			sum = sum*10 + b
			break
		} else {
			sum = sum*10 + b
		}
		x= x/10
	}

	if sum >= (1 << 31) - 1    || sum <=  - (1<<31)  {return 0}
	return sum
}