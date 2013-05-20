package main

import (
  "fmt"
  "math/big"
)

func isPrime(in int) bool {
  for i:=2;i*i<=in;i++{
    if in%i==0{
      return false
    }
  }
  return true
}

func allPrimes(lim int) []int{
  primes := []int{}
  for i:=1;i<=lim;i++{
    if isPrime(i){
      primes = append(primes, i)
    }
  }
  return primes
}

func bigPow(x,y int) big.Int{
  xbig := big.NewInt(int64(x))
  res := big.NewInt(1)
  for i:=1;i<=y;i++{
    res.Mul(res,xbig)
  }
  return *res
}

func main(){
  lim := 1000
  primes := allPrimes(lim)
  mx_cyc := 0
  mx_len := 0

  for _,i := range primes {
    j := 1
    for j<i{
      p := new(big.Int)
      *p = bigPow(10,j)
      res := new(big.Int)
      res.Mod(p,big.NewInt(int64(i)))
      if res.Int64() == 1{
        if mx_len < j{
          mx_len = i
          mx_cyc = j
        }
        break
      }
      j++
    }
  }
  fmt.Printf("Found: 1/%d yields %d cycles\n", mx_len, mx_cyc)
}
