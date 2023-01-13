package main

import "fmt"

// 阶乘
func main() {
	result := 0
	result = factorial(15)
	fmt.Println(result)
}

func factorial(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = n * factorial(n-1)
	}
	return res
}
