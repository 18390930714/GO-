<h2>strconv包</h2>
Go语言中strconv包实现了基本数据类型和其字符串表示的相互转换。<br>
strconv包实现了基本数据类型与其字符串表示的转换，主要有以下常用函数： Atoi()、Itoa()、parse系列、format系列、append系列。<br>

<h4>string与int类型转换</h4>

<h5>Atoi( )</h5>
Atoi()函数用于将字符串类型的整数转换为int类型，函数签名如下。<br>
func Atoi(s string) (i int, err error)<br>
如果传入的字符串参数无法转换为int类型，就会返回错误。<br>

<h5>Itoa()</h5>
Itoa()函数用于将int类型数据转换为对应的字符串表示，具体的函数签名如下。<br>
func Itoa(i int) string<br>

<h4>Parse系列函数</h4>
Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。<br>

<h5>ParseBool( )</h5>
func ParseBool(str string) (value bool, err error) <br>
返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。<br>

<h5>ParseInt( )</h5>
func ParseInt(s string, base int, bitSize int) (i int64, err error)<br>
返回字符串表示的整数值，接受正负号。<br>

base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；<br>

bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；<br>

返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。<br>

<h5>ParseUnit()</h5>
func ParseUint(s string, base int, bitSize int) (n uint64, err error)<br>

ParseUint类似ParseInt但不接受正负号，用于无符号整型。<br>

<h5>ParseFloat</h5>
解析一个表示浮点数的字符串并返回其值。<br>

如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。<br>

bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；<br>

返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。<br>

<h5>Format系列函数</h5>
Format系列函数实现了将给定类型数据格式化为string类型数据的功能。<br>

<h5>FormatBool()</h5>
func FormatBool(b bool) string <br>
根据b的值返回”true”或”false”。<br>

<h5>FormatInt()</h5>
func FormatInt(i int64, base int) string<br>
返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。<br>

<h5>FormatUint()</h5>
func FormatUint(i uint64, base int) string <br>
是FormatInt的无符号整数版本。<br>

<h5>FormatFloat()</h5>
func FormatFloat(f float64, fmt byte, prec, bitSize int) string <br>
函数将浮点数表示为字符串并返回。<br>
bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。<br>


<h5>isPrint()</h5>
func IsPrint(r rune) bool <br>
返回一个字符是否是可打印的，和unicode.IsPrint一样，r必须是：字母（广义）、数字、标点、符号、ASCII空格。<br>

<h5>CanBackquote()</h5>
func CanBackquote(s string) bool<br>
返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串。<br>