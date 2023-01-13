package main

import "fmt"

//使用递归函数从10打印到1

func main() {
	fmt.Println(INTprint(10))
}

func INTprint(n int) (res int) {
	fmt.Println(n)
	if n > 2 {
		return INTprint(n - 1)
	} else {
		return 1
	}
}
