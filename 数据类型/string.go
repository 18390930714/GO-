package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "开发技术"
	world := "努力学习jjjjj"

	ss := name + world //拼接字符串

	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world) // Sprintf() 返回拼接后的字符串
	fmt.Printf("ss1 = %s\n", ss1)
	fmt.Printf("%s%s", name, world)
	s3 := `D:\workspace\webDeploy` // ` ` 反引号表示不加修饰，直接输出字符串

	fmt.Printf("s3 type : %T\n", s3) // ` ` 输出类型为string

	ret := strings.Split(s3, "\\") //字符切片  以 \ 为分隔符切片 得到结果为: [D: workspace webDeploy]
	fmt.Println(ret)
	fmt.Printf("ret type is %T\n", ret) // 结果为 []string 字符串数组

	fmt.Println(strings.Join(ret, "+")) // 将字符用 ”+“ 号拼接

	fmt.Println(strings.Contains(ss, "理想"))  // 包含: 字符串中是否包含 "理想" 字符, 输出 bool值 false | true
	fmt.Println(strings.HasPrefix(ss, "理想")) // 判断前缀: 前缀是否为 "理想"，输出bool值
	fmt.Println(strings.HasSuffix(ss, "理想")) // 判断后缀: 后缀是否为 ”理想“, 输出bool值 

	s4 := "absdfb"

	fmt.Println(strings.Index(s4, "c"))     // 字符串索引 输出索引标号
	fmt.Println(strings.LastIndex(s4, "b")) // 获取某个字符最后出现的位置

}
