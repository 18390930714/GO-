package main

import "fmt"

func main(){
  /* rune 是类型别名,常用来表示 汉语、韩语、阿拉伯语等文字
     1、uint8类型也就是byte型, 代表了ASCII码的一个字符
     2、rune类型代表一个UTF-8编码字符也就是int32
  */
  s2 := "白萝卜"
  s3 := []rune(s2)  //将字符串强制转换成一个rune切片 => ['白' '萝' '卜']
  fmt.Printf("s3 type is %T\n", s3) // rune切片的字符类型为 []int32
  s3[0] = '红'
  fmt.Println(string(s3))

  c1 := "红" // 双引号表示字符串
  c2 := '红' // 单引号表示字符
  fmt.Printf("c1:%T c2:%T\n", c1, c2) // c1是字符串类型为string; c2为int32因为rune类型为int32
  //c3 := "H"
  c3 := "H"  // string
  c4 := 'H'  // int32
  fmt.Printf("c3:%T c4:%T\n", c3, c4)
}
