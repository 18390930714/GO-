package main

import "fmt"


  // 数组
  // 存放元素的容器
  // 必须指定存放的元素类型和容量(长度)
  // 数组的长度是数组类型的一部分 eg: a1长度为3 a2长度为4 则a1与a2不是同一个数据类型

func main() {
  //声明数组
  var a1 [3]bool
  var a2 [4]bool

  fmt.Printf("a1:%T a2:%T\n", a1, a2) //不初始化,默认元素都是零值(布尔值：false, 整型和浮点型都是0，字符串："")

  //初始化方式一
  a1 = [3]bool{true, true, true}
  fmt.Println(a1)  //[true, true, true]
  //初始化方式二: 根据初始值自动推断数组长度
  a3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
  fmt.Println(a3)  //[0, 1, 2, 3, 4, 5, 6, 7]
  //初始化方式三: 根据索引来初始化
  a4 := [5]int{1,2}
  fmt.Println(a4)  //[1, 2, 0, 0, 0]

  // 数组的遍历

  citys := [...]string{"北京", "上海", "深圳"}

  // for 循环遍历
  for i := 0; i < len(citys); i++ {
     fmt.Println(citys[i])
  }

  //for range 遍历
  for i, v := range citys {
     fmt.Println(i, v)
  }

  // 多维数组 [3][2]int  三行两列的数组
}
