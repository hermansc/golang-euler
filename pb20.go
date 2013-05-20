package main

import (
  "fmt"
  "math/big"
  "strconv"
)

func factorial(num int) *big.Int{
  res := new(big.Int)
  res = big.NewInt(1)
  for i:=num;i>0;i--{
    res.Mul(res, big.NewInt(int64(i)))
  }
  return res
}

func main(){
  res := new(big.Int)
  res = factorial(100)
  sum := 0
  for _,v := range res.String(){
    val,_ := strconv.Atoi(string(v))
    sum += val
  }
  fmt.Println(sum)
}
