package main

import "fmt"

func main(){
  //浮点数 float32 \ float64
  // math.MaxFloat32  float32的最大值
  f1 := 1.23456
  fmt.Printf("%T\n",f1) // float64  GO语言默认小数都是float64类型
  f2 := float32(1.234565)
  fmt.Printf("%T\n", f2) // error: f1=f2, float32类型的值不能赋值给float64类型的变量,反之也不行
}
