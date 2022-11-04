package main

import "fmt"

// 定义基于int类型的自定义类型
type Myint int

// 类型别名
type NewInt = int

func main() {
	var a Myint
	fmt.Printf("a tpye is: %T，a value is：%v\n", a, a)

	var b NewInt
	fmt.Printf("b type is：%T, a  value is：%v\n", b, b)
}
