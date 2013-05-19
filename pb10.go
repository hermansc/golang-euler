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

func main(){
  sum := 0
  lim := 2000000
  for i:=2;i<lim;i++{
    if isPrime(i){
      sum += i
    }
  }
  fmt.Println(sum)
}
