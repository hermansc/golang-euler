package main

import (
  "fmt"
)

func getNext(n int, seq []int) []int {
  if n < 2{
    return append(seq, n)
  } else {
    if n%2==0{
      r := getNext(n/2, seq)
      return append(r, n)
    } else {
      r := getNext(3*n + 1, seq)
      return append(r, n)
    }
  }
}

func getSequence(n int) []int{
  seq := []int{}
  seq = getNext(n, seq)
  return seq
}

func main(){
  longestChainLength := 0
  longestNum := 0
  lim := 1000000
  for i:=0;i<lim;i++{
    sequence := getSequence(i)
    if len(sequence) > longestChainLength{
      longestNum = i
      longestChainLength = len(sequence)
    }
  }
  fmt.Printf("Number %d produces a chain of length %d\n", longestNum, longestChainLength)
}
