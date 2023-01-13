package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

// output: ret : 2
