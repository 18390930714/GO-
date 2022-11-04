package main

import "fmt"

func main() {
  var a int = 100
  var b int = 300

  fmt.Printf("a内存地址为 %x\n", &a)
  fmt.Printf("b内存地址为 %x\n", &b)
  fmt.Printf("before, a = %d\n", a)
  fmt.Printf("before, b = %d\n", b)
  
  swap(&a,&b)

  fmt.Printf("last, a = %d\n", a)
  fmt.Printf("last, b = %d\n", b)
  fmt.Printf("a内存地址为 %x\n", &a)
  fmt.Printf("b内存地址为 %x\n", &b)
}

func swap( x *int, y *int){
  var temp int
  temp = *x
  *x = *y
  *y = temp
}
