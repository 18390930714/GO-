package main

import "log"

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[吴彦祖]") // 日志设置前缀
	log.Println("这是一条很高级的日志。")
}

/*执行后输出：
2022/11/01 15:29:11.417221 D:/go/log包/log03.go:7: 这是一条很普通的日志。
[吴彦祖]2022/11/01 15:29:11.445783 D:/go/log包/log03.go:9: 这是一条很高级的日志。
*/
