package main

import (
	"fmt"
	"reflect"
)

/*
结构体反射：
任意通过reflect.TypeOf()
*/
type studenta struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

func main() {
	stu1 := studenta{
		Name:  "白石洲",
		Score: 59,
	}

	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("name: %v type: %v tag: %v\n", f.Name, f.Type, f.Tag)
		//	fmt.Println(f.Tag.Get(key: "json"), f.Tag.Get(key: "ini"))
	}
	f2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v\n", f2.Name, f2.Type, f2.Tag)
	}
}
