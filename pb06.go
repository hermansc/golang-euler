package main

import (
  "fmt"
  "math"
)

func main(){
  sumOfSquares := 0.0
  squareOfSum := 0.0
  lim := 100
  for i:=1;i<=lim;i++{
    sumOfSquares += math.Pow(float64(i),2)
    squareOfSum += float64(i)
  }
  squareOfSum = math.Pow(float64(squareOfSum),2)

  diff := squareOfSum - sumOfSquares
  fmt.Println(diff)
}
