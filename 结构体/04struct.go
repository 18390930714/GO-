package main

import "fmt"

// 结构体方法继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s 会动\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //匿名嵌套, 且嵌套的是一个函数体指针
}

func (d *Dog) wang() {
	fmt.Printf("%s 会汪汪叫\n", d.name)
}

func main() {
	d1 := Dog{
		Feet: 4,
		Animal: &Animal{ // 取地址符&
			name: "大黄狗",
		},
	}
	d1.wang()
	d1.move()
}
