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
  a_start := -999
  b_start := -999
  n_start := 0
  a_limit := 1000
  b_limit := 1000
  num_primes := 0
  sum_coef_maxPrimes := 0
  coefs := ""
  for a:=a_start;a<a_limit;a++{
    for b:=b_start;b<b_limit;b++{
      for n:=n_start;;n++{
        res := n*n + (a)*n + b
        if res > 0{
          if !isPrime(res){
            if n > num_primes{
              num_primes = n
              sum_coef_maxPrimes = a * b
              coefs = fmt.Sprintf("(a: %d, b: %d)", a, b)
            }
            break
          }
        } else {
          break
        }
      }
    }
  }
  fmt.Printf("Found a consecutive %d primes with sum %d %s\n", num_primes, sum_coef_maxPrimes, coefs)
}
