package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作
func main() {
	//打开文件
	fileObj, err := os.Open("./test.txt") //open函数应用只读场景, ./test.txt表示当前路径
	fmt.Printf("%T\n", fileObj)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// defer 当函数退出时及时关闭file句柄,否则会内存泄露
	defer fileObj.Close()
	// 读取文件
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Printf("read from file failed, err:%v\n", err)
		return
	}
	fmt.Printf("read %d bytes from file. \n", n)
	fmt.Println(string(tmp))
}

// 循环读取大文件
func ReadAll() {
	fileObj, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("open all file  failed, err:%v\n", err)
	}
	defer fileObj.Close()

	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF { //io.EOF: End Of File
			//把当前读了多少个字节数据打印出来然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err: %v\n", n)
			return
		}
		fmt.Printf("read %d bytes from file.\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

func readByBufio() {
	fileObj, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("read file by bufio failed, err: %v\n", err)
		return
	}
	fmt.Print(line)
}

// ioutil 快速的读取文件
func readByIoutil() {
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("read file by iioutil failed, err: %v\n", err)
		return
	}
	fmt.println(string(content))
}
