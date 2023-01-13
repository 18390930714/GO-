package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4

type Stacks struct {
	index int
	data  [LIMIT]int
}

func main() {
	st1 := new(Stacks)
	fmt.Printf("%v\n", st1)
	fmt.Println("1***********************")
	st1.Push(3)
	fmt.Printf("%v\n", st1)
	fmt.Println("2***********************")
	st1.Push(7)
	fmt.Printf("%v\n", st1)
	fmt.Println("3***********************")
	st1.Push(10)
	fmt.Printf("%v\n", st1)
	fmt.Println("4***********************")
	st1.Push(99)
	fmt.Printf("%v\n", st1)
	fmt.Println("5***********************")
	p := st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Println("6***********************")
	fmt.Printf("%v\n", st1)
	fmt.Println("7***********************")
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Println("8***********************")
	fmt.Printf("%v\n", st1)
	fmt.Println("9***********************")
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Println("10***********************")
	fmt.Printf("%v\n", st1)
	fmt.Println("11***********************")
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Println("12***********************")
	fmt.Printf("%v\n", st1)
	fmt.Println("13***********************")
}

func (s *Stacks) Push(n int) {
	if s.index+1 > LIMIT {
		return
	}
	s.data[s.index] = n
	s.index++
}

func (s *Stacks) Pop() int {
	s.index--
	return s.data[s.index]
}

func (s *Stacks) String() string {
	str := ""
	for i := 0; i < s.index; i++ {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
