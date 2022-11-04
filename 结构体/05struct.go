package main

import (
	"encoding/json"
	"fmt"
)

/* 结构体字段的可见性和json序列化
# 如果一个GO语言包中定义的标识符是首字母大写的，那么就是对外可见的
# 如果一个结构体中的字段名首字母是大写的，那么该字段就是对外可见的
type student struct {
  id int  小写开头的字段对其他包不可见，其他包无法调用
  Name string
}
#结构体tag：Tag 是结构体的元信息，可以在运行的时候通过反射机制读取出来，
type class struct {
  Title string
}
*/

type student struct {
	ID   int
	Name string
}

// 构造函数
func createStudent(id int, name string) student {
	return student{
		id,
		name,
	}
}

type class struct {
	Title    string
	Students []student
}

func main() {
	// 创建一个班级变量c1
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20), // slice 类型变量要初始化
	}
	// 往班级c1中添加学生
	for i := 0; i < 10; i++ {
		tmpStu := createStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}

	fmt.Printf("%#v\n", c1)
	// JSON序列化：GO语言数据类型 转换成 JSON序列

	data, err := json.Marshal(c1)
	if err != nil { //报错打印
		fmt.Println("json marshal faild , err: ", err)
		return
	}
	fmt.Printf("data数据类型：%T\n", data) // []uint8 字节类型 byte,所以JSON数据类型为字节类型
	fmt.Printf("data数据值：%s\n", data)
	//JSON反序列化
	jsonStr := `{"Title":"长征5B","Students":[{"ID":901,"Name":"cz001"},{"ID":911,"Name":"cz002"},{"ID":930,"Name":"cz003"}]}` // `` 反斜单引号标记的数据表示直接输出的字符串
	fmt.Printf("jsonStr的数据类型：%T\n", jsonStr)
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2) // []byte(jsonStr) 表示将字符串jsonStr转换成字节类型,JSON 数据就是byte 字节类型 = []uint8
	if err != nil {
		fmt.Println("json unmarshal failed, err: ", err)
		return
	}
	fmt.Printf("%#v\n", c2)
}
