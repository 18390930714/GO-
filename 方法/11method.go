package main

import (
	"fmt"
	"strconv"
)

/*
实现栈（stack）数据结构：
push 将一个新值放到栈的最顶部一个非空（非零）的格子中。

pop 获取栈的最顶部一个非空（非零）的格子的值。现在可以理解为什么栈是一个后进先出（LIFO）的结构了吧。

为栈定义一 Stack 类型，并为它定义一个 Push 和 Pop 方法，再为它定义 String() 方法（用于调试）它输出栈的内容，比如：[0:i] [1:j] [2:k] [3:l]。

1）stack_arr.go：使用长度为 4 的 int 数组作为底层数据结构。

*/

const LIMIT = 4

type Stack [LIMIT]int

func main() {
	st1 := new(Stack)
	fmt.Printf("%v\n", st1)
}

func (s Stack) Push(n int) { // push方法将一个数加入栈的顶部
	for i, v := range s {
		if v == 0 {
			s[i] = n
			break
		}
	}
}

func (s Stack) Pop() int { //pop获取顶部第一个非空栈的值
	v := 0
	for i := len(s); i >= 0; i-- {
		if v = s[i]; v != 0 {
			s[i] = 0
			return v
		}
	}
	return 0
}

func (s Stack) string() string {
	str := ""
	for i, v := range s {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(v) + "]"
	}
	return str
}
