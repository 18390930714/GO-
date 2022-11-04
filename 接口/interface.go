package main

import "fmt"

/*
一、GO语言中接口是一种类型，一种抽象的类型
二、值接收者实现接口与使用指针接收者实现接口的区别
三、一个结构体实现多个接口
四、接口的嵌套
五、空接口: 接口中没有定义任何方法时，任意类型都实现了空接口 -->空接口变量可以存储任意值
六、类型断言
*/

type Phone interface {
  call()
}

type NokiaPhone struct {
}

// 方法实现接口
func (nokiaPhone NokiaPhone) call() {
  fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iphone IPhone) call() {
  fmt.Println("I am IPhone, I will call you!")
}

func main() {
  var mynokiaphone Phone
  var myiphone Phone

  mynokiaphone = new(NokiaPhone)
  mynokiaphone.call()

  myiphone = new(IPhone)
  myiphone.call()

  var x interface{} //定义一个空接口变量
  x = "hello"
  fmt.Println(x)
  x = 100
  fmt.Println(x)
  x = false
  fmt.Println(x)
  var m = make(map[string]interface{}, 16)
  m["name"] = "阿迪力"
  m["age"] = 18
  m["hobby"] = []string{"篮球", "游泳", "游戏"}
  fmt.Println(m)

  //类型断言
  ret, ok := x.(string) // 猜对了ok返回值是true，猜错了ok返回值是false
  if !ok {
    fmt.Println("不是字符串")
  } else {
    fmt.Println("是字符串类型", ret)
  }
}
