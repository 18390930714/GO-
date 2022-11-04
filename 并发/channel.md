<h2>channel通道</h2>
如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。<br>
Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。<br>
每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。<br>

<h4>channel类型</h4>
channel是 Go 语言中一种特有的类型。声明通道类型变量的格式如下：<br>
var 变量名称 chan 元素类型<br>

其中：<br>
chan：是关键字<br>
元素类型：是指通道中传递元素的类型<br>

举几个例子：<br>
var ch1 chan int   // 声明一个传递整型的通道<br>
var ch2 chan bool  // 声明一个传递布尔型的通道<br>
var ch3 chan []int // 声明一个传递int切片的通道<br>

<h5>初始化channel</h5>
声明的通道类型变量需要使用内置的make函数初始化之后才能使用。具体格式如下：<br>
make(chan 元素类型, [缓冲大小])<br>
其中：<br>
channel的缓冲大小是可选的。<br>
举几个例子：<br>

ch4 := make(chan int)<br>
ch5 := make(chan bool, 1)  // 声明一个缓冲区大小为1的通道<br>

<h5>channel操作</h5>
通道共有发送（send）、接收(receive）和关闭（close）三种操作。而发送和接收操作都使用<-符号。<br>
现在我们先使用以下语句定义一个通道：<br>

ch := make(chan int)<br>
发送<br>
将一个值发送到通道中。<br>

ch <- 10 // 把10发送到ch中<br>
接收<br>
从一个通道中接收值。<br>

x := <- ch // 从ch中接收值并赋值给变量x<br>
<-ch       // 从ch中接收值，忽略结果<br>
关闭<br>
我们通过调用内置的close函数来关闭通道。<br>

close(ch)<br>

<h5>无缓冲的通道</h5>
无缓冲的通道又称为阻塞的通道。我们来看一下如下代码片段。<br>
func main() {<br>
&emsp;&emsp;ch := make(chan int)<br>
&emsp;&emsp;ch <- 10<br>
&emsp;&emsp;fmt.Println("发送成功")<br>
}<br>
上面这段代码能够通过编译，但是执行的时候会出现以下错误：<br>

fatal error: all goroutines are asleep - deadlock!<br>

goroutine 1 [chan send]:<br>
main.main()<br>
&emsp;&emsp;&emsp;&emsp;.../main.go:8 +0x54<br>
因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有接收方能够<br>
接收值的时候才能发送成功，否则会一直处于等待发送的阶段。同理，如果对一个无缓冲通道执行接收操作时，<br>
没有任何向通道中发送值的操作那么也会导致接收操作阻塞。就像田径比赛中的4x100接力赛，想要完成交<br>
棒必须有一个能够接棒的运动员，否则只能等待。简单来说就是无缓冲的通道必须有至少一个接收方才能发送成功。<br>

<h5>有缓冲通道</h5>
还有另外一种解决上面死锁问题的方法，那就是使用有缓冲区的通道。我们可以在使用 make 函数初始化通道时，可以为其指定通道的容量，例如：<br>

func main() {<br>
&emsp;&emsp;ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道<br>
&emsp;&emsp;ch <- 10<br>
&emsp;&emsp;fmt.Println("发送成功")<br>
}<br>
只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量。当通道内已有元素数达到最大容量后， <br>
再向通道执行发送操作就会阻塞，除非有从通道执行接收操作。就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，<br>
等到别人取走一个快递员就能往里面放一个。<br>

我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。<br>

<h4>多返回值</h4>
<h4>单向通道</h4>
<h4>select多路复用</h4>