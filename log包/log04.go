package main

import "log"

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[吴彦祖]")
	log.Println("这是一条很普通的日志。")
}
