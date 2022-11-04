package main

import (
	"fmt"
	"reflect"
)

/* GO语言反射
反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。
GO语言的反射中的数组、切片、map 指针等类型的变量，它们的返回值都是空.
*/

func reflectType(x interface{}) {
	// 调用本函数时不知道会传入什么类型的参数
	//1. 方式一：通过类型断言
	//2. 方式二：借助反射
	v := reflect.TypeOf(x) //获取变量类型信息
	fmt.Println(v)
	fmt.Println(v, v.Name(), v.Kind())
	fmt.Printf("%T\n", v)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v, %T\n", v, v)
}

type Cat struct{}

type Dog struct{}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
	var c Cat
	reflectType(c)
	var d Dog
	reflectType(d)
}
