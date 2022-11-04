package main

import (
	"fmt"
)

//方法定义

// Person 结构体
type Person struct {
	name string
	age  int8
}

// 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream() 是为Person类型定义的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好GO语言！\n", p.name)
}

// 修改年龄的方法
func (p *Person) setAge(newAge int8) { // p 是指针类型的接收者
	p.age = newAge
}

func main() {
	p1 := NewPerson("zzz", int8(19)) // NewPerson函数返回值是指针
	//(*p1).Dream()  理解上是需要这样调用方法，取指针所指向的值
	p1.Dream()

	//调用修改年龄的方法
	fmt.Println(p1.age)
	p1.setAge(int8(29))
	fmt.Println(p1.age)
}
