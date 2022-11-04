package main

import (
	"fmt"
)

//嵌套结构体

type Address struct {
	Province string
	City     string
}

type Person struct {
	Name    string
	Gender  string
	Age     int
	Address Address //嵌套另一个结构体
}

// 嵌套匿名结构体
type Person2 struct {
	Name    string
	Gender  string
	Age     int
	Address //嵌套另一个匿名结构体
}

// 嵌套结构体字段冲突
type Email struct {
	Username string
	City     string
}

type Person3 struct { //Address字段与Email字段都有元素City，冲突;
	Name   string
	Gender string
	Age    int
	Address
	Email
}

func main() {
	p1 := Person{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province: "湖南",
			City:     "长沙",
		},
	}

	p2 := Person2{
		Name:   "大王子",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province: "广东",
			City:     "广州",
		},
	}

	p3 := Person3{
		Name:   "三王子",
		Gender: "男",
		Age:    19,
		Address: Address{
			Province: "湖南",
			City:     "岳阳",
		},
		Email: Email{
			Username: "123456",
			City:     "衡阳",
		},
	}

	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Name, p1.Gender, p1.Age, p1.Address)
	fmt.Println(p2.Address.Province) //通过嵌套的匿名结构体字段访问其内部字段
	fmt.Println(p2.Province)         //直接访问匿名结构体中的字段
	// fmt.Println(p3.City) 报错： .\02struct.go:81:17: ambiguous selector p3.City 识别不出是哪个City
	fmt.Println(p3.Address.City)
	fmt.Println(p3.Email.City)
}
