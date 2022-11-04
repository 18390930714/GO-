package main

import (
	"fmt"
	"time"
)

// time demo
func main() {
	now := time.Now() // Now函数返回一个时间对象
	fmt.Println(now)
	year := now.Year()
	mouth := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, mouth, day, hour, minute, second)
	//将时间戳转换为具体的时间格式
	n := 5
	// Duration 时间间隔类型;
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("over")
	fmt.Println("Add 函数:")
	fmt.Println(now)
	//Add
	t2 := now.Add(time.Hour)
	fmt.Println(t2)
	// Sub
	fmt.Println(t2.Sub(now))
	//定时器
	time.Tick(time.Second)
	//时间格式化
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("20060102150405"))

	//将字符串时间解析成go语言格式时间
}
