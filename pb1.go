package main

import (
  "fmt"
)

func findSum(lim int) int {
  sum := 0
  for i:=0;i<lim;i++ {
    if i%3==0 || i%5==0 {
      sum += i
    }
  }
  return sum
}

func main() {
  sum := findSum(1000)
  fmt.Println(sum)
}
