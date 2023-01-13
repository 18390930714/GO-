package main

import (
	"fmt"
)

func main() {
	var f = plus()
	fmt.Print(f(1), "-")
	fmt.Print(f(20), "-")
	fmt.Print(f(300))
}

func plus() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

/*
三次调用函数f过程中函数plus()中的变量delta的值分别是1\20\300
多次调用中，变量x的值被保留的，即0 + 1 = 1 , 然后1 + 20 = 21 , 最后21 + 300 = 321 ; 闭包函数保存并积累其中的变量的值
不管外部函数是否退出都能够继续操作外部函数中的局部变量
*/
