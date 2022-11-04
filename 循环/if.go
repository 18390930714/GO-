package main

import "fmt"

func main(){
  /* 
   作用域：age 变量此时只在if条件判断语句中生效
  */
  if age := 19; age > 18 {    //生成一个age的局部变量 只在if 判断中生效
    fmt.Printf("我 %d 岁 了\n", age)
  } else {
    fmt.Println("我 还 是 个 未 成 年！")
  }
}
