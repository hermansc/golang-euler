package main

import (
  "fmt"
)

func findPrimes2(num int) []int{
  primes := []int{}
  for i:=2;i<num;i++{
    if num%i==0{
      primes = append(primes,i)
      num = num/i
      i = 1
    }
  }
  primes = append(primes, num)
  return primes
}

func Max(a []int) int {
  max := 0
  for _,v := range a {
    if v>max{
      max = v
    }
  }
  return max
}

func main() {
  primes := findPrimes2(600851475143)
  maxVal := Max(primes)
  fmt.Println(maxVal)
}
