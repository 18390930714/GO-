package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张山", "姓名")
	flag.IntVar(&age, "age", 19, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟时间间隔")
	flag.Parse() //解析命令行参数
	fmt.Println(name, age, married, delay)
	fmt.Println("*****************")
	fmt.Println(flag.Args()) //返回命令行参数后的其他参数，以[]string类型
	fmt.Println("*****************")
	fmt.Println(flag.NArg()) //返回命令行参数后的其他参数个数
	fmt.Println("*****************")
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数
}
