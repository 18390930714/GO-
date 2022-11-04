package split

import "strings"

// Split 将s按照sep进行分割，返回一个字符串的切片
// Split("我爱你",”爱“) => ["我"，"你"]

func Split(s, sep string) (ret []string) {
	idx := strings.Index(s, sep) // Index 函数 获取s 字符串中 sep 的位置索引
	for idx > -1 {
		ret = append(ret, s[:idx]) // append函数 向ret字符串切片中追加一个切片
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
