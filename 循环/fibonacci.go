package main

import "fmt"

func main() {
 var i int
 for i = 0; i < 11; i++ {
  fmt.Printf("%d\t", fibonacci(i))
 }

}

func fibonacci(n int) int {
  if n < 2 {
    return n
  }
  return fibonacci(n-2) + fibonacci(n-1)
}
