package main

import (
	"fmt"
)

// 将函数作为参数
func main() {
	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("the sum if %d an %d is %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}
