package main

import "fmt"

func main() {
	//声明切片类型
	var a []string              //声明一个字符切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              // []
	fmt.Println(b)              // []
	fmt.Println(c)              // [false, true]
	fmt.Println(d)              // [false, true]
	//初始化
	s1 := []int{1, 2, 3}
	s2 := []string{"沙河", "张江", "平山"}
	//长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 2.由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13} //底层数组
	s3 := a1[0:4]                         //基于一个数组切割,左包含右不包含,（左闭右开）
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4] // => [0:4] [1 3 5 7]
	s6 := a1[3:] // => [3:len(a1)] [7 9 11 13]
	s7 := a1[:]  // => {0:len(a1)}
	fmt.Println(s5, s6, s7)
	//1.切片指向了一个底层数组
	//2.切片的长度就是元素个数
	//3.切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	//底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))
	//切片的切片
	s8 := s6[3:] //[]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))
}
