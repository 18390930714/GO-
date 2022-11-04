package main

import "fmt"

func main() {
  i1 := 19
  fmt.Printf("%d\n", i1)  // %d: 显示十进制
  fmt.Printf("%b\n", i1)  // 把十进制转换成二进制 %b:显示二进制
  fmt.Printf("%o\n", i1)  // 把十进制转换成八进制 %o:显示八进制
  fmt.Printf("%x\n", i1)  // 把十进制转换成十六进制 %x : 显示十六进制

  
  i2 := 077  //定义八进制数
  fmt.Printf("%d\n", i2)  // 八进制转换成十进制， %d 显示十进制

  i3 := 0x123 //定义十六进制数
  fmt.Printf("%d\n", i3) // 十六进制转十进制

  fmt.Printf("%T\n", i3) // %T 查看变量数据类型


}
