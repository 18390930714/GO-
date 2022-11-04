package main

import "log"

func main() {
	//在记录日志之前先设置标准logger的输出选项
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}
