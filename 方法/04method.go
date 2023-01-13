package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time //将类型作为匿名类型嵌入在一个新的结构体中
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}

	fmt.Println("Full time now", m.String())
	fmt.Println("First 3 chars:", m.first3Chars())
}
