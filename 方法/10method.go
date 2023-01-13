package main

import "fmt"

/*
为 int 定义别名类型 TZ，定义一些常量表示时区，比如 UTC，定义一个 map，它将时区的缩写映射为它的全称，
比如：UTC -> "Universal Greenwich time"。为类型 TZ 定义 String() 方法，它输出时区的全称。
*/

const (
	UTC TZ = iota
	EST
	CST
)

var timeZones = map[TZ]string{
	UTC: "Universal Greenwich time",
	EST: "Eastern Standard time",
	CST: "Central Standard time"}

func (tz TZ) String() string {
	if zone, ok := timeZones[tz]; ok {
		return zone
	}
	return ""
}

func main() {
	fmt.Println(EST)
	fmt.Println(UTC)
	fmt.Println(CST)
}
