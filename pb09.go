package main

import (
  "fmt"
  "math"
)

func pythagoreanTriplet(lim float64) int{
  for a:=1;a<int(lim)/2;a++{
    for b:=a+1;b<int(lim)/2;b++{
      af := float64(a)
      bf := float64(b)
      cexp := math.Pow(af,2) + math.Pow(bf,2)
      c := math.Sqrt(cexp)
      if af+bf+c == lim && bf < c {
        return int(af*bf*c)
      }
    }
  }
  return 0
}

func main(){
  sum := pythagoreanTriplet(1000.0)
  fmt.Println(sum)
}
