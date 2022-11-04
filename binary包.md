<h1>binary包</h1>
<h2>字节序</h2>
<h3>一、字节序定义</h3>
<p>字节序就是多字节数据类型(int, float)等在内存或通信链路中的存储顺序。在网络传输中基于文本类型的协议(比如JSON)<br>
和二进制协议都是字节通信，采用字节序进行数据包的处理。</p>
<p>字节序可以分位大端序：内存低位地址存放数据的高位字节；小端序相反：内存低位地址存放数据低位字节。<br>
![](images/18473-5932c0bfcc2abac9.png)<br>
在计算机内部，小端序被广泛应用于现代性 CPU 内部存储数据；而在其他场景譬如网络传输和文件存储使用大端序。在网络协议层<br>
操作二进制数字时约定使用大端序，大端序是网络字节传输采用的方式。因为大端序最高有效字节排在首位（低地址端存放高位字节）,<br>
能够按照字典排序，所以我们能够比较二进制编码后数字的每个字节。

<h2>固定长的编码</h2>
Go 中有多种类型的整型， int8, int16, int32 和 int64 ，分别使用 1, 3, 4, 8 个字节表示，我们称之为固定长度类型 (fixed-length types)。<br>

<h3>Go处理固定长度字节</h3>

Go中处理大小端序的代码位于 encoding/binary ,包中的全局变量BigEndian用于操作大端序数据，LittleEndian用于操作小端序数据，<br>这两个变量所对应的数据类型都实行了ByteOrder接口：<br>
binary包实现了对数据与byte之间的转换，以及varint的编解码。
type ByteOrder interface { <br>
    Uint16([]byte) uint16  <br>
    Uint32([]byte) uint32  <br>
    Uint64([]byte) uint64  <br>
    PutUint16([]byte, uint16)  <br>
    PutUint32([]byte, uint32)  <br>
    PutUint64([]byte, uint64)  <br>
    String() string            <br>
}                              <br>
    
其中，前三个方法用于读取数据，后三个方法用于写入数据。

<h3>Go处理固定长度流</h3>
binary package 提供了内置的读写固定长度值的流。<br>
func Read(r io.Reader, order ByteOrder, data interface{}) error  <br>
1、r可以读出字节流的数据源<br>
2、order 特殊字节，包中提供的大端字节序和小端字节序<br>
3、data 需要解码成的数据<br>
返回值：error  返回错误<br>
功能说明：Read从r中读出字节数据并反序列化成结构数据。data必须是固定长的数据值或固定长数据的slice。从r中读出的数据可以使用特殊的 字节序来解码，并顺序写入value的字段。当填充结构体时，使用(_)名的字段讲被跳过。<br>
func Write(w io.Writer, order ByteOrder, data interface{}) error  <br>
1）w  可写入字节流的数据<br>
2）order  特殊字节序，包中提供大端字节序和小端字节序<br>
3）data  需要解码的数据<br>
返回值：error  返回错误<br>
功能说明：
Write讲data序列化成字节流写入w中。data必须是固定长度的数据值或固定长数据的slice，或指向此类数据的指针。写入w的字节流可用特殊的字节序来编码。另外，结构体中的(_)名的字段讲忽略。<br>