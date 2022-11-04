package main

import "fmt"

func main() {
  var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
  for  key, value :=  range pow {       //range 的返回值为key value 键值对，for _, value range 或者  for key range 这样缩写
    fmt.Printf("2**%d == %d\n", key, value)
  }
}
