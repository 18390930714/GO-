package main

import (
	"fmt"
	"strconv"
)

type T struct {
	a int
	b float32
	c string
}

/*
练习 10.12 type_string.go

给定结构体类型 T:

	type T struct {
	    a int
	    b float32
	    c string
	}

值 t: t := &T{7, -2.35, "abc\tdef"}。给 T 定义 String()，使得 fmt.Printf("%v\n", t) 输出：7 / -2.350000 / "abc\tdef"。

eg: 输出：7 / -2.350000 / "abc\tdef"
*/
func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("t is : %v\n", t)
}

func (t *T) String() string {
	return strconv.Itoa(t.a) + "/" + strconv.FormatFloat(float64(t.b), 'f', 6, 32) + "/" + string(t.c)
}
