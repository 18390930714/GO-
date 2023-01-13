type name struct {} 结构体定义<br>

<h3>使用new</h3>
使用new函数给一个新的结构体分配内存，返回指向已分配内存的指针：var t *T = new(T)<br>
t := new(T)  变量t事一个指向T的指针，此时结构体字段的值是它们所属类型的零值;<br>

type myStruct struct {i int} <br>
var v myStruct  //v是结构体类型变量 <br>
var p *myStruct  //p 是指向一个结构体类型变量的指针 <br>
v.i <br>
p.i <br>
无论变量是一个结构体类型还是一个结构体类型指针，都使用同样的 选择器符（selector-notation） 来引用结构体的字段: v.i 与 p.i的值相同<br>

初始化一个结构体实列的更简短和惯用的方式如下：
ms := &struct1{10, 15.5, "Chris"} , ms是指向结构体的指针；
混合字面量语法 &struct{a,b,c} 是一种简写，底层仍会调用new(),表达式new(Type) 和&Type{}是等价的。

<h3>使用工厂方法创建结构体实列</h3>
GO语言不支持面向对象编程语言那样的构造方法：但可以实现类似的结构体方法：

type File struct { <br>
    fd int   //文件描述符 <br>
    name string  //文件名  <br>
} <br>

对应的工厂方法：<br>
func NewFile(fd int, name string) *File {<br>
&emsp;&emsp;&emsp;&emsp;if fd < 0 {<br>
&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;return nil<br>
}<br>
&emsp;&emsp;&emsp;&emsp;return &File{fd, name}<br>
}<br>

调用工厂方法：<br>
f := NewFile(10, "./test.txt")<br>

<h3>使用自定义包中的结构体</h3>

<h3>带标签的结构体</h3>
标签tag是附属于字段的字符串，struct的tag只有reflect包才能获取，具体在反射章中描述；

<h3>匿名字段和内嵌结构体</h3>
匿名字段没有显式的名字，只有字段的类型是必须的；<br>
eg: 08struct.go <br>

