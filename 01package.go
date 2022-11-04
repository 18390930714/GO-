package main

// GO语言中不允许导入包而不使用！！
// GO语言中不允许循环引用包！！

// 当代码都放在$GOPATH目录下,导入包的路径要从$GOPATH/src后面继续写起

import (
	bazha "code/video18/package_demo/calc" // 包的别名
	"fmt"
)

func main() {
	fmt.Println(bazha.Name)
}
