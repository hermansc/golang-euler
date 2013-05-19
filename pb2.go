package main

import (
  "fmt"
)

var cache map[int]int

func fibonacci(i int) int {
  if _, ok := cache[i]; !ok {
    if (i < 3){
      cache[i] = i
    } else{
      res := fibonacci(i-1) + fibonacci(i-2)
      cache[i] = res
    }
  }
  return cache[i]
}

func main() {
  cache = make(map[int]int)
  sum := 0
  lim := 4000000
  for i:=0;i<lim;i++{
    s := fibonacci(i)
    if s>lim{
      break
    } else if s % 2 == 0 {
      sum += s
    }
  }
  fmt.Println(sum)
}
