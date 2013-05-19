package main

import (
  "fmt"
  "math/big"
  "strconv"
)

func main(){
  res := big.NewInt(1)
  exp := 1000
  for i:=0;i<exp;i++{
    res.Mul(res,big.NewInt(2))
  }
  sum := 0
  for _,v := range res.String(){
    val,_ := strconv.Atoi(string(v))
    sum += val
  }
  fmt.Println(sum)
}
