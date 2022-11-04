package main

import "fmt"

const ip = 3.148

/*批量常量*/
const (
  statusOk = 200
  notFound = 404
)

/*常量未赋值则以上一行常量值为默认值*/
const (
  n1 = 100
  n2
  n3
)

/*iota*/

const (
  a1 = iota  // a1 = 0
  a2         // a2 = 1
  a3         // a3 = 2

)

/* 使用_跳过某些值*/
const (
  b1 = iota // b1 = 0 , iota = 0
  b2        // b2 = 1 , iota = 1
  _         // 匿名变量：_ = iota ,被丢弃 iota = 2
  b3        // b3 = 3 , iota = 3
)

/* ioat 声明中间插队*/
const (
  c1 = iota // c1 = iota = 0
  c2 = 100  // c2 = 100, iota = 1
  c3        // c3 = 100
)

const (
  c11 = iota // c11 = iota = 0
  c21 = 100  // c21 = 100 , iota =1
  c31 = iota // c31 = iota = 2
  c41        // c41 = ioat = 3
)
/* iota 多个常量在一行声明*/
const (
  d1, d2 = iota+1, iota+2 // iota=0, d1 = iota + 1 = 0 + 1 =1, d2 = iota + 2 = 0 + 2 =2
  d3, d4 = iota+1, iota+2 // iota=1, d3 = iota + 1 = 1 + 1 =2, d4 = iota + 2 = 1 + 2 =3
)

/*定义数量级*/
const (
  _  = iota
  KB = 1 << (10 * iota) // 二进制, 1 << (10 * 1) , 1 左移 10位, 等于二进制 10000000000 = 十进制 1024
  MB = 1 << (10 * iota) // 二进制, 1 << (10 * 2) , 1 左移 20位, 等于二进制 100000000000000000000
  GB = 1 << (10 * iota)
  TB = 1 << (10 * iota)
  PB = 1 << (10 * iota)
)

func main() {
  fmt.Println("n1:", n1)
  fmt.Println("n2:", n2)
  fmt.Println("n3:", n3)

  fmt.Println("c1:", c1)
  fmt.Println("c2:", c2)
  fmt.Println("c3:", c3)

  fmt.Println("c11:", c11)
  fmt.Println("c21:", c21)
  fmt.Println("c31:", c31)
  fmt.Println("c41:", c41)

  fmt.Println("d1:", d1)
  fmt.Println("d2:", d2)
  fmt.Println("d3:", d3)
  fmt.Println("d4:", d4)

}
