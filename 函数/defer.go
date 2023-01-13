package main

import "fmt"

// defer是推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
func a() {
	i := 0
	defer fmt.Println(i) // defer接收参数
	i++
	return
}

func main() {
	a()
}
