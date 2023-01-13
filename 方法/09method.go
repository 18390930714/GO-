package main

import "fmt"

/*为 int 定义一个别名类型 Day，定义一个字符串数组它包含一周七天的名字，为类型 Day 定义 String() 方法，
它输出星期几的名字。使用 iota 定义一个枚举常量用于表示一周的中每天（MO、TU...）。
*/

type Day int

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

const (
	MO Day = iota // Day类型的常量MO = iota = 0; iota 可以看作一个自增索引
	TU            // TU = 1
	WE            // WE = 2
	TH            // TH = 3
	FR            // FR = 4
	SA            // SA = 5
	SU            // SU = 6
)

func main() {
	var th Day = 3
	fmt.Printf("the 3rd day is : %s\n", th)

	var day = SU
	fmt.Println(day)
	fmt.Println(0, MO, 1, TU)
}

func (d Day) String() string {
	return week[d]
}
