package main

import (
  "fmt"
)

func smallestMultiple(lim int) int {
  for i:=10;;i++{
    for j:=1;j<=lim;j++{
      if i%j == 0{
        if j == lim{
          return i
        }
      } else {
        break
      }
    }
  }
}

func main() {
  s := smallestMultiple(20)
  fmt.Println(s)
}
