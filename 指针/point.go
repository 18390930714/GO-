package main

import "fmt"

/*
& 取地址符
* 根据地址取值
指针传值
*/
func main() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("&a 取出a的地址为: %x\n", &a)
	fmt.Printf("ip = %x\n", ip)
	fmt.Printf("指针ip取a的值为: %d\n", *ip)
	/*
				new() make()
				var a *int //声明了一个空指针 a = nil
				*a = 100    //空指针未分配内存地址，会报错，需要用new函数初始化
				new() make() 都是初始化函数，都是为了占用一个内存空间
				new() :  func new(Type) *Type ; Type表示类型 输入数据类型 输出类型指针
			    make():  只用于slice、map、channel的初始化
		        new 用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
	*/

}
