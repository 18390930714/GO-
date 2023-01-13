package main

import (
	"fmt"
	"strconv"
)

/*
为 float64 定义一个别名类型 Celsius，并给它定义 String()，它输出一个十进制数和 °C 表示的温度值。
*/

type Celsius float64

func main() {
	var c1 Celsius
	c1 = 5.67899
	fmt.Printf("c1 is: %v\n", c1)
}

func (c Celsius) String() string {
	return strconv.Itoa(int(c)) + "°C"
}
