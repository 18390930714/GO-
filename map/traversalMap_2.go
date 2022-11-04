package main

import "fmt"

func main() {
  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    fmt.Printf("num = %d\n", num)
    sum += num
  }
  fmt.Printf("sum: %d\n", sum)

  for i, num := range nums { 
    if num == 3 {
      fmt.Printf("key: %d\n", i)
    }
  }

  kvs :=  map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
