package main

import "fmt"

type IntVector []int //切片

// 非结构体类型上方法
func main() {
	fmt.Println(IntVector{1, 2, 3}.sum())
}

func (v IntVector) sum() (s int) { //切片的方法
	for _, x := range v {
		s += x
	}
	return s
}
