package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args是一个[]string 字符串切片
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

/*flag.Type() 方法: flag.Type(flag名, 默认值, 帮助信息)*Type
name := flag.String("name", "张三", "姓名")
age := flag.Int("age", 18, "年龄")
married := flag.Bool("married", false, "婚否")
delay := flag.Duration("d", 0, "时间间隔")
需要注意的是，此时name、age、married、delay均为对应类型的指针。
*/

/*
flag,TypeVar() 方法：flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
var name string
var age int
var married boll
var delay time.Duration
flag.StringVar(&name, name, "张三", "姓名")
flag.IntVar(&age, "age", 18, "年龄")
flag.BoolVar(&married, "married", false, "婚否")
flag.DurationVar(&delay, "d", 0, "时间间隔")
*/

/*flag.Parse()
通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。

支持的命令行参数格式有以下几种：
-flag xxx （使用空格，一个-符号）
--flag xxx （使用空格，两个-符号）
-flag=xxx （使用等号，一个-符号）
--flag=xxx （使用等号，两个-符号）
其中，布尔类型的参数必须使用等号的方式指定。
Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。*/

/* flag其他函数
flag.Args()  ////返回命令行参数后的其他参数，以[]string类型
flag.NArg()  //返回命令行参数后的其他参数个数
flag.NFlag() //返回使用的命令行参数个数
*/
