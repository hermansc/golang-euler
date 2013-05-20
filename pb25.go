package main

import (
  "fmt"
  "math/big"
)

var cache map[int]big.Int

func fibonacci(i int) big.Int {
  if _, ok := cache[i]; !ok {
    res := new(big.Int)
    k := new(big.Int)
    j := new(big.Int)
    *k = fibonacci(i-1)
    *j = fibonacci(i-2)
    res.Add(k,j)
    cache[i] = *res
  }
  return cache[i]
}

func main(){
  cache = make(map[int]big.Int)
  cache[0] = *big.NewInt(0)
  cache[1] = *big.NewInt(1)
  for i:=2000;;i++{
    s := new(big.Int)
    *s = fibonacci(i)
    if len(s.String()) == 1000{
      fmt.Println(i)
      break
    }
  }
}
