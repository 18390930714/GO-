package main

import "fmt"

/*
channel 是goroutine并发的函数之间传递数据的通道
channel 是一种数据类型、一种引用类型;
channel 通道遵循先进先出的原则

声明通道：
var ch1 chan int
var ch2 chan bool
var ch3 chan []int

初始化通道  make(chan 元素类型, [缓冲大小])  缓存大小表示通道可存的值的数量
ch4 := make(chan int)  无缓冲区通道 又称同步通道
ch5 := make(chan int, 1) 有缓冲区的通道 又称伊异步通道

channel 操作
发送: 将值发送到通道中: ch <- 10
接收: 从通道中接收值： x := <- ch 以及  <- ch 忽略结果
关闭: 通道的关闭不是必须的： 建议关闭：close(ch)
*/
func main() {
	var ch1 chan int
	ch1 = make(chan int, 1) //通道是引用类型，需要make初始化
	ch1 <- 10               //发送值到通道
	x := <-ch1              //接受通道ch1的值
	fmt.Println(x)          //打印值
	//len(ch1)                //取通道中的元素数
	//cap(ch1)                //取通道的容量
	close(ch1)

	ch2 := make(chan int) //无缓冲区通道
	ch2 <- 20
	y := <-ch2
	fmt.Println(y)

}
