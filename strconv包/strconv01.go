package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "100"
	i1, err := strconv.Atoi(s1) //Itoa()函数用于将int类型数据转换为对应的字符串表示
	if err != nil {
		fmt.Println("cant convert to int")
	} else {
		fmt.Printf("type: %T, value: %v\n", i1, i1)
	}
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type: %T, value: %#v\n", s2, s2)
}
