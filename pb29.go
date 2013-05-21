package main

import (
  "fmt"
  "math/big"
)

func removeDuplicates(xs *[]*big.Int){
  found := make(map[string]bool)
  j := 0
  for i,x := range *xs{
    if !found[x.String()]{
      found[x.String()] = true
      (*xs)[j] = (*xs)[i]
      j++
    }
  }
  *xs = (*xs)[:j]
}

func bigPow(x,y int) *big.Int{
  xbig := big.NewInt(int64(x))
  res := big.NewInt(1)
  for i:=1;i<=y;i++{
    res.Mul(res,xbig)
  }
  return res
}

func main(){
  sums := []*big.Int{}
  for a:=2;a<=100;a++{
    for b:=2;b<=100;b++{
      powed := bigPow(a,b)
      sums = append(sums, powed)
    }
  }
  removeDuplicates(&sums)
  fmt.Println(len(sums))
}
