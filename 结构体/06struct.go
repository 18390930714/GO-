package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

//初始化结构体的三种方法

func main() {
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	pers2 := new(Person) // pers2 是指针类型
	pers2.firstName = "rick"
	(*pers2).lastName = "Max"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	pers3 := &Person{"Bob", "jackson"} // pers3 是指向结构体的指针
	upPerson(pers3)
	fmt.Printf("the name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}
