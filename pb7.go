package main

import (
  "fmt"
)

func isPrime(in int) bool {
  for i:=2;i*i<=in;i++{
    if in%i==0{
      return false
    }
  }
  return true
}

func findPrime(lim int) int {
  for i:=0;;i++ {
    if isPrime(i){
      if lim < 0{
        return i
      }
      lim--
    }
  }
}

func main() {
  num := findPrime(10001)
  fmt.Println(num)
}
