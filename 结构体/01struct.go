package main

import "fmt"

//结构体的匿名字段: 结构体的元素没有名字

type Person struct { // type 声明的都是类型
	string
	int
}

func main() {
	p1 := Person{
		string: "小王子",
		int:    18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string, p1.int)
}
