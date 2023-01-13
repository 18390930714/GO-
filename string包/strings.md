<h2>前缀和后缀</h2>
前缀：strings.HasPrefix(s, prefix string) bool
后缀：strings.HasSuffix(s, suffix string) bool

<h2>字符串包含关系</h2>
strings.Contains(s, substr string) bool

<h2>判断字符串或字符在父字符串中出现的位置 (索引)</h2>
strings.Index(s, str string) int

返回字符串在父字符串中最后一次出现的位置的索引：strings.LastIndex(s, str string) int

strings.IndexRune(s string, r rune) int

<h2>字符串替换</h2>
strings.Replace(str, old, new string, n int) string

<h2>统计字符串出现次数</h2>
strings.Count(s, str string) int

<h2>重复字符串</h2>
strings.Repeat(s, count int) string

<h2>修改字符串大小写</h2>
strings.ToUpper(s) string
strings.ToLower(s) string

<h2>修剪字符串</h2>
strings.TrimSpace(s) 剔除开头和结尾的空白符号；
strings.Trim(s, "cut") 将字符串s中开头和结尾的cut去除掉
strings.TrimLeft() 剔除开头
strings.TrimRight() 剔除结尾

<h2>分割字符串</h2>
strings.Fields(s) 利用空白作为分隔符将字符串分割为若干块，并返回一个 slice 
strings.Split(s, sep) 自定义分隔符号对字符串分割，返回slice.

<h2>拼接slice到字符串</h2>
strings.Join(s1 []string, sep string) string

<h2>从字符串中读取内容</h2>
strings.NewReader(str) 用于生成一个 Reader 并读取字符串中的内容，然后返回指向该 Reader 的指针;




