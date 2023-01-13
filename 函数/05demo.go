package main

import "fmt"

// 闭包应用：将函数作为返回值
func main() {
	p2 := Add() // p2 是Add函数的返回值，Add函数的返回值是个函数
	fmt.Printf("The Add for 3 gives: %v\n", p2(3))

	p3 := Adder(2)
	fmt.Printf("The result is : %v\n", p3(3))
}

func Add() func(b int) int {
	return func(b int) int {
		return b + 3
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
