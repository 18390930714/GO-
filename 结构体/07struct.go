package main

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	z := make(Bar) // 编译错误：cannot make type Bar
	(*z).thingOne = "hello"
	(*z).thingTwo = 1

	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "poor life"

	u := new(Foo)
	(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	(*u)["y"] = "world"
}
