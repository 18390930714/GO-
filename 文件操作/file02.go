package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 写入文件

func main() {
	write()
}

func write() {
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//标志位：os.O_CREATE 文件不存在则创建，os.O_WRONLY：只写标志位，os.O_APPEND：写入文件标志位
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	str := "沙河小王子"
	fileObj.Write([]byte(str))           //接收字节类型
	fileObj.WriteString("hello junfeng") //接收字符串类型
}

func writeByBufio() {
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//标志位：os.O_CREATE 文件不存在则创建，os.O_WRONLY：只写标志位，os.O_APPEND：写入文件标志位
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("迪丽热巴") //将内容写入缓冲区
	writer.Flush()             //将缓冲区内容写入磁盘
}

func writeByioutil() {
	str := "深圳赚钱深圳花一分别想带回家"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed  err: %v\n", err)
	}
}
