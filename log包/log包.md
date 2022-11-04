<h2>log</h2>
Go语言内置的log包实现了简单的日志服务。本文介绍了标准库log的基本使用。

<h3>使用logger</h3>
log包定义了Logger类型，该类型提供了一些格式化输出的方法。本包也提供了一个预定义的“标准”logger，<br>
可以通过调用函数Print系列(Print|Printf|Println）、Fatal系列（Fatal|Fatalf|Fatalln）、和Panic系列（Panic|Panicf|Panicln）来使用，<br>
比自行创建一个logger对象更容易使用。

<h3>配置logger</h3>
<h4>标准logger的配置</h4>

默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。<br>
log标准库中为我们提供了定制这些设置的方法。<br>
log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置。<br>

func Flags() int<br>
func SetFlags(flag int)<br>

<h4>flag 选项</h4>
log标准库提供了如下的flag选项，它们是一系列定义好的常量。<br>

const (<br>
// &emsp;控制输出日志信息的细节，不能控制输出的顺序和格式。<br>
// &emsp;输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message <br>
Ldate       &emsp;  = 1 << iota  &emsp;   // 日期：2009/01/23<br>
Ltime                 &emsp;    &emsp;    // 时间：01:23:23<br>
Lmicroseconds       &emsp;&emsp;          // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）<br>
Llongfile             &emsp;&emsp;        // 文件全路径名+行号： /a/b/c/d.go:23<br>
Lshortfile             &emsp;&emsp;       // 文件名+行号：d.go:23（会覆盖掉Llongfile）<br>
LUTC                   &emsp;&emsp;       // 使用UTC时间<br>
LstdFlags  &emsp;   = Ldate | Ltime &emsp;// 标准logger的初始值<br>
)

<h4>配置日志前缀</h4>
log标准库中还提供了关于日志信息前缀的两个方法：<br>

func Prefix() string<br>
func SetPrefix(prefix string)<br>

<h4>配置日志输出位置</h4>
func SetOutput(w io.Writer)<br>
SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。<br>

<h4>创建新logger</h4>


log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。New函数的签名如下：<br>

func New(out io.Writer, prefix string, flag int) *Logger<br>